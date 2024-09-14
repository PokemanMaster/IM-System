package utils

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	Red *redis.Client
)

// InitConfig 引入 viper
func InitConfig() {
	viper.SetConfigName("app")    // 配置文件的名称为 app，viper 将会寻找一个名为 app.yml、app.json 等的文件作为配置文件。
	viper.AddConfigPath("config") // 指定 viper 在哪个目录下查找配置文件。在 config 目录下查找配置文件。
	err := viper.ReadInConfig()   // 读取并加载配置文件。 viper.ReadInConfig() 会在之前指定的目录中寻找配置文件并读取其内容。
	if err != nil {
		fmt.Println("读取配置文件错误!", err)
	}
	fmt.Println("config app.yml init")
}

// InitMySQL 初始化 MySQL
func InitMySQL() {
	//自定义日志模板 打印SQL语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL阈值
			LogLevel:      logger.Info, //级别
			Colorful:      true,        //彩色
		},
	)

	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")),
		&gorm.Config{Logger: newLogger})

	fmt.Println("MySQL init")
}

// InitRedis 初始化 Redis
func InitRedis() {
	Red = redis.NewClient(&redis.Options{
		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})
	pong, err := Red.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("Init Redis err", err)
	} else {
		fmt.Println("Init Redis", pong)
	}
}

const (
	PublishKey = "websocket"
)

// Publish 发布消息到Redis
func Publish(ctx context.Context, channel string, msg string) error {
	var err error
	fmt.Println("Publish", msg)
	err = Red.Publish(ctx, channel, msg).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

// Subscribe 订阅Redis消息
func Subscribe(ctx context.Context, channel string) (string, error) {
	sub := Red.Subscribe(ctx, channel)
	fmt.Println("Subscribe", ctx)
	msg, err := sub.ReceiveMessage(ctx)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("Subscribe", msg.Payload)
	return msg.Payload, err
}
