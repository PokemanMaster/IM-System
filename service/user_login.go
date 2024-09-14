package service

import (
	"IMProject/models"
	"IMProject/pkg/e"
	"IMProject/serializer"
	"IMProject/utils"
	"github.com/gin-gonic/gin"
)

type UserLoginService struct {
	Name     string
	Password string
}

// UserLogin
// @Summary 所有用户
// @Tags 用户模块
// @param name query string false "用户名"
// @param password query string false "密码"
// @Success 200 {string} json{"code","message"}
// @Router /user/findUserByNameAndPwd [post]
func (service *UserLoginService) UserLogin(c *gin.Context) *serializer.Response {
	data := models.UserBasic{}
	code := e.SUCCESS
	name := service.Name
	password := service.Password

	user := models.FindUserByName(name)
	if user.Name == "" {
		code = e.ERROR_MATCHED_USERNAME
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	flag := utils.ValidPassword(password, user.Salt, user.PassWord)
	if !flag {
		code = e.ERROR_MATCHED_USERNAME
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}

	pwd := utils.MakePassword(password, user.Salt)
	data = models.FindUserByNameAndPwd(name, pwd)

	return &serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   data,
	}
}
