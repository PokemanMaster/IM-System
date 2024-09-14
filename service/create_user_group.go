package service

import (
	"IMProject/models"
	"IMProject/pkg/e"
	"IMProject/serializer"
	"IMProject/utils"
	"github.com/gin-gonic/gin"
)

type CreateCommunityService struct {
	OwnerId uint
	Name    string
	Icon    string
	Desc    string
}

// CreateCommunity 新建群
func (service *CreateCommunityService) CreateCommunity(c *gin.Context) *serializer.Response {
	ownerId := service.OwnerId
	name := service.Name
	icon := service.Icon
	desc := service.Desc

	community := models.Community{}
	community.OwnerId = ownerId
	community.Name = name
	community.Img = icon
	community.Desc = desc

	code, msg := models.CreateCommunity(community)
	if code == 0 {
		utils.RespOK(c.Writer, code, msg)
	} else {
		utils.RespFail(c.Writer, msg)
	}

	cod := e.SUCCESS
	return &serializer.Response{
		Status: code,
		Msg:    e.GetMsg(cod),
	}
}
