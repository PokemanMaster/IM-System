package api

import (
	"IMProject/pkg/logging"
	"IMProject/service"
	"github.com/gin-gonic/gin"
)

// RedisMsg 获取用户A、B的消息
func RedisMsg(ctx *gin.Context) {
	services := service.RedisMsgService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, ErrorResponse(err))
		logging.Info(err)
	} else {
		res := services.RedisMsg()
		ctx.JSON(200, res) // 解析数据JSON
	}
}

//// SendUserMsg 发送用户消息
//func SendUserMsg(ctx *gin.Context) {
//	services := service.SendUserMsgService{}
//	err := ctx.ShouldBind(&services)
//	if err != nil {
//		ctx.JSON(400, ErrorResponse(err))
//		logging.Info(err)
//	} else {
//		res := services.Chat(ctx)
//		ctx.JSON(200, res) // 解析数据JSON
//	}
//}

// Upload 上传文件
func Upload(ctx *gin.Context) {
	services := service.AddFriendService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, ErrorResponse(err))
		logging.Info(err)
	} else {
		res := services.AddFriend(ctx)
		ctx.JSON(200, res) // 解析数据JSON
	}
}

// SendMsg 发送消息
func SendMsg(ctx *gin.Context) {
	services := service.AddFriendService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, ErrorResponse(err))
		logging.Info(err)
	} else {
		res := services.AddFriend(ctx)
		ctx.JSON(200, res) // 解析数据JSON
	}
}
