package models

import (
	"IMProject/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/spf13/viper"
	"gopkg.in/fatih/set.v0"
	"gorm.io/gorm"
)

// Message 消息
type Message struct {
	gorm.Model
	UserId     int64  //发送者
	TargetId   int64  //接受者
	Type       int    //发送类型  1私聊  2群聊  3心跳
	Media      int    //消息类型  1文字 2表情包 3语音 4图片 /表情包
	Content    string //消息内容
	CreateTime uint64 //创建时间
	ReadTime   uint64 //读取时间
	Pic        string
	Url        string
	Desc       string
	Amount     int //其他数字统计
}

func (table *Message) TableName() string {
	return "message"
}

var clientMap = make(map[int64]*Node, 0) // ClientMap 映射关系
var rwLocker sync.RWMutex                // RwLocker 读写锁

// Node 节点
type Node struct {
	Conn          *websocket.Conn //连接
	Addr          string          //客户端地址
	FirstTime     uint64          //首次连接时间
	HeartbeatTime uint64          //心跳时间
	LoginTime     uint64          //登录时间
	DataQueue     chan []byte     //消息
	GroupSets     set.Interface   //好友 / 群
}

// Heartbeat 更新用户心跳
func (node *Node) Heartbeat(currentTime uint64) {
	node.HeartbeatTime = currentTime
	return
}

// sendProc(node *Node) - 发送消息处理逻辑
func sendProc(node *Node) {
	for { // 确保程序持续运行以处理消息发送。
		select { // 用于监听多路通道，这里只监听 node.DataQueue 通道。
		case data := <-node.DataQueue: // 当 DataQueue 通道中有数据时，接收数据并存储在 data 变量中。
			// 这里发送的是 TextMessage 类型（WebSocket 文本消息），数据内容为 data。
			err := node.Conn.WriteMessage(websocket.TextMessage, data) // 通过 WebSocket 连接发送消息。
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

// recvProc(node *Node) - 接收消息处理逻辑
func recvProc(node *Node) {
	for { // 持续监听 WebSocket 连接上的消息。
		_, data, err := node.Conn.ReadMessage() // 通过 WebSocket 连接读取消息，data 是接收到的消息内容。
		if err != nil {                         // 如果读取消息时发生错误（如连接断开），打印错误并退出循环。
			fmt.Println("node.Conn err:", err)
			return
		}
		msg := Message{}                 // 声明一个 Message 类型的变量 msg，用于存储解码后的消息。
		err = json.Unmarshal(data, &msg) // 将接收到的 JSON 数据解码为 msg 结构体。
		if err != nil {                  // 如果解码失败，打印错误。
			fmt.Println("JSON 解析失败：", err)
		}

		// 心跳检测 msg.Media == -1 || msg.Type == 3
		if msg.Type == 3 {
			currentTime := uint64(time.Now().Unix())
			node.Heartbeat(currentTime) // 更新该节点的心跳时间。
		} else {
			dispatch(data) // 函数处理接收到的消息
			broadMsg(data) // 函数将消息广播到局域网
		}
	}
}

var udpsendChan = make(chan []byte, 1024) // udp 发送通道
func broadMsg(data []byte)                { udpsendChan <- data } // 把数据发送到udp通道中

func init() {
	go udpSendProc()
	go udpRecvProc()
	fmt.Println("init goroutine")
}

// udpSendProc - UDP 数据发送协程
func udpSendProc() {
	con, err := net.DialUDP("udp", nil, &net.UDPAddr{ // 该函数用于建立一个 UDP 连接，net.UDPAddr 指定了目标地址和端口号。
		IP:   net.IPv4(192, 168, 0, 255), // 设置广播 IP 地址（192.168.0.255），用于将数据广播到局域网内的所有主机。
		Port: viper.GetInt("port.udp"),   // 从配置文件中获取端口号（使用 viper 进行配置管理）。
	}) // con：表示通过 DialUDP 创建的 UDP 连接。
	defer con.Close()
	if err != nil { // 在连接建立时发生错误，输出错误信息并终止后续操作。
		fmt.Println(err)
	}

	// 数据发送逻辑
	for { // 监听 udpsendChan 通道。
		select {
		case data := <-udpsendChan:
			_, err := con.Write(data) // 通道中有数据时，读取数据并通过 con.Write(data) 向目标地址发送 UDP 数据包。
			if err != nil {           // 发送过程中出错，打印错误并返回退出。
				fmt.Println(err)
				return
			}
		}
	}
}

// udpRecvProc - UDP 数据接收协程
func udpRecvProc() {
	con, err := net.ListenUDP("udp", &net.UDPAddr{ //  该函数用于监听指定的 UDP 地址和端口，等待接收数据。
		IP:   net.IPv4zero,             // 表示监听所有可用网络接口的 IP 地址。
		Port: viper.GetInt("port.udp"), // 从配置中获取监听的 UDP 端口号。
	}) // con：表示创建的 UDP 监听连接。
	if err != nil {
		fmt.Println(err) // 创建监听连接时发生错误，打印错误信息并终止操作。
	}
	defer con.Close() // 使用 defer 确保在函数结束时关闭 UDP 监听连接。

	// 数据接收逻辑
	for {
		var buf [512]byte           // 使用一个大小为 512 字节的缓冲区 buf 来存储接收到的 UDP 数据。
		n, err := con.Read(buf[0:]) // 调用 con.Read(buf[0:]) 读取数据到缓冲区，返回读取的字节数 n。
		if err != nil {
			fmt.Println(err)
			return
		}
		dispatch(buf[0:n]) // 调用 dispatch 函数对接收到的 buf 数据进行处理
	}
}

// Chat 需要 ：发送者ID ，接受者ID ，消息类型，发送的内容，发送类型
func Chat(writer http.ResponseWriter, request *http.Request) {
	// 解析 URL 查询参数
	query := request.URL.Query()              // 获取请求中的 URL 查询参数。
	Id := query.Get("userId")                 // 从查询参数中获取 userId，即当前用户的 ID。
	userId, _ := strconv.ParseInt(Id, 10, 64) // 将字符串形式的 userId 转换为 int64 类型。

	// 创建 WebSocket 连接
	isvalida := true                   // 定义一个布尔变量用于 WebSocket 连接的安全校验（这里简单设为 true，允许所有请求通过）。
	conn, err := (&websocket.Upgrader{ // 升级 HTTP 请求为 WebSocket 连接。
		CheckOrigin: func(r *http.Request) bool { // CheckOrigin 是用于检查请求来源的函数，在此简单地返回 true，表示允许所有跨域请求。
			return isvalida
		},
	}).Upgrade(writer, request, nil) // 将 HTTP 请求升级为 WebSocket 连接。
	if err != nil { // 如果升级失败（err != nil），则打印错误并返回。
		fmt.Println(err)
		return
	}

	// 创建 Node 节点
	currentTime := uint64(time.Now().Unix()) // 获取当前时间戳（秒级），用于表示心跳和登录时间。
	node := &Node{                           // 创建一个 Node 结构体实例，代表当前连接的客户端信息：
		Conn:          conn,                       // 当前客户端的 WebSocket 连接。
		Addr:          conn.RemoteAddr().String(), // 客户端地址
		HeartbeatTime: currentTime,                // 心跳时间
		LoginTime:     currentTime,                // 登录时间
		DataQueue:     make(chan []byte, 50),      // 消息队列 一个用于存储待发送消息的缓冲通道（容量为 50）。
		GroupSets:     set.New(set.ThreadSafe),    // 用户组集合（线程安全） 一个线程安全的集合，用来存储当前用户加入的群组。
	}

	// 将用户 ID 与 Node 绑定
	rwLocker.Lock()
	clientMap[userId] = node // 将 userId 和 node 绑定存储在 clientMap 中，代表这个用户与他的连接状态。
	rwLocker.Unlock()        // 使用读写锁 rwLocker 确保对共享资源的安全访问。

	// 发送消息的处理逻辑
	go sendProc(node)

	// 接收消息的处理逻辑
	go recvProc(node)

	// 记录在线用户到 Redis 缓存
	// SetUserOnlineInfo：调用一个函数将当前用户标记为在线，具体逻辑是将用户的在线状态存储到 Redis 缓存中：
	// online_ + Id：Redis 的键名，表示这个用户的在线状态。
	// []byte(node.Addr)：将用户的地址存储为缓存的值。
	// time.Duration(viper.GetInt("timeout.RedisOnlineTime")) * time.Hour：用户在线状态的过期时间，从配置中读取超时时间（单位为小时）。
	SetUserOnlineInfo("online_"+Id, []byte(node.Addr), time.Duration(viper.GetInt("timeout.RedisOnlineTime"))*time.Hour)
}

// 后端调度逻辑处理
func dispatch(data []byte) {
	msg := Message{}                           // 存储解析后的消息
	msg.CreateTime = uint64(time.Now().Unix()) // 设置为当前 Unix 时间戳，以记录消息创建的时间。
	err := json.Unmarshal(data, &msg)          // 将字节数组 data 解析为 msg 对象。
	if err != nil {
		fmt.Println(err)
		return
	}
	switch msg.Type { // 根据其值执行不同的处理逻辑。
	case 1: //私信
		sendMsg(msg.TargetId, data) // 发送的好友ID ，消息内容
	case 2: //群发
		sendGroupMsg(msg.TargetId, data) // 发送的群ID ，消息内容
	}
}

// 发送私信
func sendMsg(userId int64, msg []byte) {
	rwLocker.RLock()              // 这是一个读写锁，用来在读取 clientMap 时防止并发冲突。
	node, ok := clientMap[userId] // // clientMap 是一个维护用户 ID 和 Node（用户连接信息）之间映射的全局变量。
	rwLocker.RUnlock()

	jsonMsg := Message{}
	err := json.Unmarshal(msg, &jsonMsg) // 将字节数组 msg 反序列化为 Message 结构体。
	if err != nil {
		fmt.Println("反序列化为 Message 结构体err", err)
		return
	}

	// 检查用户是否在线:
	ctx := context.Background()
	targetIdStr := strconv.Itoa(int(userId))
	userIdStr := strconv.Itoa(int(jsonMsg.UserId))
	jsonMsg.CreateTime = uint64(time.Now().Unix())
	r, err := utils.Red.Get(ctx, "online_"+userIdStr).Result() // 使用 utils.Red.Get 从 Redis 中获取用户的在线状态（online_userId）。
	if err != nil {
		fmt.Println(err)
	}

	// 发送消息到在线用户:
	if r != "" {
		if ok { // 如果用户在线并且在 clientMap 中找到对应的 Node，则将消息通过 node.DataQueue 发送给用户。
			node.DataQueue <- msg
		}
	}

	// 构建 Redis 键:
	var key string               // 组合出一个消息的键，用来存储消息历史。
	if userId > jsonMsg.UserId { // 消息的存储顺序取决于两个用户 ID 的大小（保证键的唯一性）。
		key = "msg_" + userIdStr + "_" + targetIdStr
	} else {
		key = "msg_" + targetIdStr + "_" + userIdStr
	}

	// 保存消息到 Redis:
	res, err := utils.Red.ZRevRange(ctx, key, 0, -1).Result() // 使用 utils.Red.ZRevRange 获取当前对话中的历史消息。
	if err != nil {
		fmt.Println(err)
	}
	score := float64(cap(res)) + 1
	// 然后使用 utils.Red.ZAdd 将消息添加到 Redis 的有序集合中，以便以后检索历史消息。
	ress, e := utils.Red.ZAdd(ctx, key, &redis.Z{score, msg}).Result() //jsonMsg
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(ress)
}

// 发送群消息
func sendGroupMsg(targetId int64, msg []byte) { // 主要功能是将消息发送给一个群组的所有成员。
	//userIds := SearchUserByGroupId(uint(targetId)) // 根据群组 ID 获取所有该群的用户 ID。
	//for i := 0; i < len(userIds); i++ {            // 遍历所有群成员
	//	if targetId != int64(userIds[i]) { // 并排除消息的发送者（即消息不发给自己）。
	//		sendMsg(int64(userIds[i]), msg) // 使用 sendMsg 函数，将消息发送给每一个群成员。
	//	}
	//}
	sendMsg(targetId, msg)
}

// MarshalBinary 需要重写此方法才能完整的msg转byte[]
func (msg Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(msg)
}

// RedisMsg 获取缓存里面的消息
func RedisMsg(userIdA int64, userIdB int64, start int64, end int64, isRev bool) []string {
	rwLocker.RLock()
	rwLocker.RUnlock()
	ctx := context.Background()
	userIdStr := strconv.Itoa(int(userIdA))
	targetIdStr := strconv.Itoa(int(userIdB))
	var key string
	if userIdA > userIdB {
		key = "msg_" + targetIdStr + "_" + userIdStr
	} else {
		key = "msg_" + userIdStr + "_" + targetIdStr
	}
	var rels []string
	var err error
	if isRev {
		rels, err = utils.Red.ZRange(ctx, key, start, end).Result()
	} else {
		rels, err = utils.Red.ZRevRange(ctx, key, start, end).Result()
	}
	if err != nil {
		fmt.Println(err) //没有找到
	}
	return rels
}

// CleanConnection 清理超时连接
func CleanConnection(param interface{}) (result bool) {
	result = true
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("cleanConnection err", r)
		}
	}()
	//node.IsHeartbeatTimeOut()
	currentTime := uint64(time.Now().Unix())
	for i := range clientMap {
		node := clientMap[i]
		if node.IsHeartbeatTimeOut(currentTime) {
			fmt.Println("心跳超时关闭连接：", node)
			node.Conn.Close()
		}
	}
	return result
}

// IsHeartbeatTimeOut 用户心跳是否超时
func (node *Node) IsHeartbeatTimeOut(currentTime uint64) (timeout bool) {
	if node.HeartbeatTime+viper.GetUint64("timeout.HeartbeatMaxTime") <= currentTime {
		fmt.Println("心跳超时自动下线：", node)
		timeout = true
	}
	return
}
