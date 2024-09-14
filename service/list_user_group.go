package service

import (
	"IMProject/models"
	"IMProject/pkg/e"
	"IMProject/serializer"
	"github.com/gin-gonic/gin"
)

type LoadCommunityService struct {
	OwnerId uint
}

// LoadCommunity 加载群列表
func (service *LoadCommunityService) LoadCommunity(c *gin.Context) *serializer.Response {
	ownerId := service.OwnerId
	data, msg := models.LoadCommunity(ownerId)
	code := e.SUCCESS

	if len(data) != 0 {
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   data,
		}
	} else {
		return &serializer.Response{
			Status: code,
			Msg:    msg,
		}
	}

}
