package api

import (
	"IMProject/pkg/logging"
	"IMProject/service"
	"github.com/gin-gonic/gin"
)

// AddFriend 添加用户
func AddFriend(ctx *gin.Context) {
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

// SearchFriends 好友列表
func SearchFriends(ctx *gin.Context) {
	services := service.FriendsService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, ErrorResponse(err))
		logging.Info(err)
	} else {
		res := services.SearchFriends(ctx)
		ctx.JSON(200, res) // 解析数据JSON
	}
}

// FindByID 查询好友
func FindByID(ctx *gin.Context) {
	services := service.FriendsService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, ErrorResponse(err))
		logging.Info(err)
	} else {
		res := services.FindByID(ctx)
		ctx.JSON(200, res) // 解析数据JSON
	}
}
