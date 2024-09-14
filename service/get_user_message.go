package service

import (
	"IMProject/models"
	"IMProject/pkg/e"
	"IMProject/serializer"
)

type RedisMsgService struct {
	UserIdA int
	UserIdB int
	Start   int
	End     int
	IsRev   bool
}

// RedisMsg 通过 Redis 获取消息记录的请求。
func (service *RedisMsgService) RedisMsg() *serializer.Response {
	userIdA := service.UserIdA
	userIdB := service.UserIdB
	start := service.Start
	end := service.End
	isRev := service.IsRev
	// 通过调用 models.RedisMsg 获取 Redis 中的消息记录
	res := models.RedisMsg(int64(userIdA), int64(userIdB), int64(start), int64(end), isRev)
	// 将查询到的消息结果返回给客户端。
	cod := e.SUCCESS
	return &serializer.Response{
		Status: cod,
		Msg:    e.GetMsg(cod),
		Data:   res,
	}
}
