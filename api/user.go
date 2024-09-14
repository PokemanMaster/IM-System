package api

import (
	"IMProject/pkg/logging"
	"IMProject/service"
	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册接口
func UserRegister(ctx *gin.Context) {
	services := service.UserRegisterService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, ErrorResponse(err))
		logging.Info(err)
	} else {
		res := services.UserRegister()
		ctx.JSON(200, res) // 解析数据JSON
	}
}

// UserLogin 用户登录接口
func UserLogin(ctx *gin.Context) {
	services := service.UserLoginService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, ErrorResponse(err))
		logging.Info(err)
	} else {
		res := services.UserLogin(ctx)
		ctx.JSON(200, res) // 解析数据JSON
	}
}

// UserUpdate 修改用户信息
func UserUpdate(ctx *gin.Context) {
	services := service.UserUpdateService{}
	err := ctx.ShouldBind(&services)
	if err != nil {
		ctx.JSON(400, ErrorResponse(err))
		logging.Info(err)
	} else {
		res := services.UserUpdate(ctx)
		ctx.JSON(200, res) // 解析数据JSON
	}
}
