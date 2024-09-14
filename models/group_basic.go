package models

import (
	"IMProject/utils"
	"gorm.io/gorm"
)

// GroupBasic 群信息
type GroupBasic struct {
	gorm.Model
	Name    string
	OwnerId uint
	Icon    string
	Type    int
	Desc    string
}

func (table *GroupBasic) TableName() string {
	return "group_basic"
}

// JoinGroup 加入群
func JoinGroup(userId uint, comId string) (int, string) {
	contact := Contact{}
	contact.OwnerId = userId
	contact.Type = 2
	community := Community{}

	utils.DB.Where("id=? or name=?", comId, comId).Find(&community)
	if community.Name == "" {
		return -1, "没有找到群"
	}
	utils.DB.Where("owner_id=? and target_id=? and type = 2 ", userId, comId).Find(&contact)
	if !contact.CreatedAt.IsZero() {
		return -1, "已加过此群"
	} else {
		contact.TargetId = community.ID
		utils.DB.Create(&contact)
		return 0, "加群成功"
	}
}
