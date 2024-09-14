package service

import (
	"IMProject/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

// 防止跨域站点伪造请求
// websocket.Upgrader 用来将 HTTP 请求升级为 WebSocket 连接。
var upGrader = websocket.Upgrader{
	// 用来检查请求的来源，防止跨站请求伪造（CSRF 攻击）。
	// 在这里，函数总是返回 true，即允许所有的跨域请求。
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// SendMsg 通过 WebSocket 发送消息的处理函数。
func SendMsg(c *gin.Context) {
	// 用来将 HTTP 请求升级为 WebSocket 协议，并建立一个 WebSocket 连接。
	// c.Writer 是 HTTP 响应的 Writer，c.Request 是 HTTP 请求对象。
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	// 如果升级失败（比如客户端不支持 WebSocket），会返回错误并终止处理。
	if err != nil {
		fmt.Println(err)
		return
	}
	// 语句确保函数退出时，WebSocket 连接会被关闭。
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(ws)
	// 成功建立 WebSocket 连接后，MsgHandler 会处理消息的发送。
	MsgHandler(c, ws)
}

// MsgHandler 持续处理消息推送的函数
func MsgHandler(c *gin.Context, ws *websocket.Conn) {
	for {
		// utils.Subscribe 订阅来自 Redis 的消息，通过 PublishKey 监听消息发布。
		msg, err := utils.Subscribe(c, utils.PublishKey)
		if err != nil {
			fmt.Println(" MsgHandler 发送失败", err)
		}
		// 获取到消息后，生成一个带有当前时间戳的字符串，并通过 WebSocket 连接推送给客户端。
		tm := time.Now().Format("2006-01-02 15:04:05")
		m := fmt.Sprintf("[ws][%s]:%s", tm, msg)
		// ws.WriteMessage(1, []byte(m)) 将消息以文本格式（1 表示文本消息类型）发送给 WebSocket 客户端。
		err = ws.WriteMessage(1, []byte(m))
		if err != nil {
			log.Fatalln(err)
		}
	}
}
