package service

import (
	"IMProject/models"
	"IMProject/pkg/e"
	"IMProject/serializer"
	"IMProject/utils"
	"github.com/gin-gonic/gin"
)

type JoinGroupsService struct {
	UserId uint
	ComId  string
}

// JoinGroups 加入群 userId uint, comId uint
func (service *JoinGroupsService) JoinGroups(c *gin.Context) *serializer.Response {
	userId := service.UserId
	comId := service.ComId
	data, msg := models.JoinGroup(userId, comId)
	if data == 0 {
		utils.RespOK(c.Writer, data, msg)
	} else {
		utils.RespFail(c.Writer, msg)
	}
	code := e.SUCCESS
	return &serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
