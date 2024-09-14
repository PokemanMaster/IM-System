package serializer

import (
	"IMProject/models"
)

// GroupSerialization 用户序列化器
type GroupSerialization struct {
	OwnerId  uint
	TargetId uint
	Type     int
	Desc     string
}

// Group 序列化用户
func Group(contact models.Contact) GroupSerialization {
	return GroupSerialization{
		OwnerId:  contact.OwnerId,
		TargetId: contact.TargetId,
		Type:     contact.Type,
		Desc:     contact.Desc,
	}
}
