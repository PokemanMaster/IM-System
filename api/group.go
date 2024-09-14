package api

import (
	"IMProject/pkg/logging"
	"IMProject/service"
	"github.com/gin-gonic/gin"
)

// ListCommunity 群好友
func ListCommunity(ctx *gin.Context) {
	services := service.ListCommunityService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, ErrorResponse(err))
		logging.Info(err)
	} else {
		res := services.List()
		ctx.JSON(200, res) // 解析数据JSON
	}
}

// CreateCommunity 用户注册接口
func CreateCommunity(ctx *gin.Context) {
	services := service.CreateCommunityService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, ErrorResponse(err))
		logging.Info(err)
	} else {
		res := services.CreateCommunity(ctx)
		ctx.JSON(200, res) // 解析数据JSON
	}
}

// LoadCommunity 用户注册接口
func LoadCommunity(ctx *gin.Context) {
	services := service.LoadCommunityService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, ErrorResponse(err))
		logging.Info(err)
	} else {
		res := services.LoadCommunity(ctx)
		ctx.JSON(200, res) // 解析数据JSON
	}
}

// JoinGroups 用户注册接口
func JoinGroups(ctx *gin.Context) {
	services := service.JoinGroupsService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, ErrorResponse(err))
		logging.Info(err)
	} else {
		res := services.JoinGroups(ctx)
		ctx.JSON(200, res) // 解析数据JSON
	}
}
