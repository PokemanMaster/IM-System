package serializer

import (
	"IMProject/models"
	"time"
)

// UserSerialization 用户序列化器
type UserSerialization struct {
	Name          string
	PassWord      string
	Phone         string
	Email         string
	Avatar        string
	Identity      string
	ClientIp      string
	ClientPort    string
	Salt          string
	LoginTime     time.Time
	HeartbeatTime time.Time
	LoginOutTime  time.Time
	IsLogout      bool
	DeviceInfo    string
}

// BuildUser 序列化用户
func BuildUser(user models.UserBasic) UserSerialization {
	return UserSerialization{

		Avatar: user.Avatar,
	}
}
