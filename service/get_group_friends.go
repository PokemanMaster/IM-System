package service

import (
	"IMProject/models"
	"IMProject/pkg/e"
	"IMProject/serializer"
	"IMProject/utils"
)

type ListCommunityService struct {
	TargetId uint
}

func (service *ListCommunityService) List() serializer.Response {
	targetId := service.TargetId

	var contact []models.Contact
	utils.DB.Where("target_id=? and type=2", targetId).Find(&contact)

	code := e.SUCCESS
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Data:   contact,
	}
}
