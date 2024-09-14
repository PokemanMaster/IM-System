package service

import (
	"IMProject/models"
	"IMProject/pkg/e"
	"IMProject/serializer"
	"IMProject/utils"
	"fmt"
	"math/rand"
	"time"
)

// UserRegisterService 前端请求过来的数据
type UserRegisterService struct {
	UserName string
	Password string
	Identity string
}

func (service *UserRegisterService) UserRegister() *serializer.Response {
	user := models.UserBasic{}
	user.Name = service.UserName
	password := service.Password
	repassword := service.Identity

	salt := fmt.Sprintf("%06d", rand.Int31())
	code := e.SUCCESS
	data := models.FindUserByName(user.Name)

	if user.Name == "" || password == "" || repassword == "" {
		code = e.ERROR_MATCHED_USERNAME
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if data.Name != "" {
		code = e.ERROR_MATCHED_USERNAME
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if password != repassword {
		code = e.ERROR_MATCHED_USERNAME
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user.PassWord = utils.MakePassword(password, salt)
	user.Salt = salt

	user.LoginTime = time.Now()
	user.LoginOutTime = time.Now()
	user.HeartbeatTime = time.Now()
	models.CreateUser(user)

	return &serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
