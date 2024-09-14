package service

import (
	"IMProject/models"
	"IMProject/pkg/e"
	"IMProject/serializer"
	"IMProject/utils"
	"github.com/gin-gonic/gin"
)

type AddFriendService struct {
	UserId     uint   // 自己的id
	TargetName string // 好友的id
}

func (service *AddFriendService) AddFriend(c *gin.Context) *serializer.Response {
	userId := service.UserId
	targetName := service.TargetName
	code, msg := models.AddFriend(userId, targetName)
	if code == 0 {
		utils.RespOK(c.Writer, code, msg)
	} else {
		utils.RespFail(c.Writer, msg)
	}

	cod := e.SUCCESS
	return &serializer.Response{
		Status: cod,
		Msg:    e.GetMsg(cod),
	}
}
