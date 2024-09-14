package main

import (
	"IMProject/models"
	"IMProject/router"
	"IMProject/utils"
	"time"

	"github.com/spf13/viper"
)

func main() {
	utils.InitConfig() // 初始化配置文件
	utils.InitMySQL()  // 初始化 MySQL
	utils.InitRedis()  // 初始化 Redis
	models.Migration()

	InitTimer()                                  // 初始化定时器
	r := router.Router()                         // 初始化路由
	panic(r.Run(viper.GetString("port.server"))) // 初始化监听端口
}

// InitTimer 初始化定时器，定时清理数据库的超时连接
func InitTimer() {
	utils.Timer(
		time.Duration(viper.GetInt("timeout.DelayHeartbeat"))*time.Second,
		time.Duration(viper.GetInt("timeout.HeartbeatHz"))*time.Second,
		models.CleanConnection, "")
}
