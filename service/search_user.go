package service

import (
	"IMProject/models"
	"IMProject/pkg/e"
	"IMProject/serializer"
	"IMProject/utils"
	"github.com/gin-gonic/gin"
)

type FriendsService struct {
	UserId uint
}

// SearchFriends 获取好友列表
func (service *FriendsService) SearchFriends(c *gin.Context) *serializer.Response {
	id := service.UserId
	users := models.SearchFriend(id)
	//utils.RespOKList(c.Writer, users, len(users))
	code := e.SUCCESS
	return &serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   users,
	}
}

// FindByID 查找某个用户
func (service *FriendsService) FindByID(c *gin.Context) *serializer.Response {
	userId := service.UserId
	data := models.FindByID(userId)
	utils.RespOK(c.Writer, data, "ok")
	code := e.SUCCESS
	return &serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}
