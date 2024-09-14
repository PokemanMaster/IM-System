package router

import (
	"IMProject/api"
	"IMProject/docs"
	"IMProject/service"
	"IMProject/utils"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(utils.CORS())
	docs.SwaggerInfo.BasePath = "" // 设置 Swagger 文档的基础路径。
	// 使用 ginSwagger.WrapHandler 将 Swagger UI 绑定到 /swagger/*any 路径，允许你通过这个路径访问 Swagger API 文档。
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	v1 := r.Group("/api/v1")
	{
		// 用户
		v1.POST("/user/register", api.UserRegister)      // 用户注册
		v1.POST("/user/login", api.UserLogin)            // 用户登录
		v1.POST("/user/deleteUser", service.DeleteUser)  // 删除用户
		v1.POST("/user/updateUser", api.UserUpdate)      // 更新用户
		v1.GET("/user/getUserList", service.GetUserList) // 用户列表

		// 群聊
		v1.POST("/contact/createCommunity", api.CreateCommunity) // 创建群聊
		v1.POST("/contact/loadcommunity", api.LoadCommunity)     // 群聊列表
		v1.POST("/contact/joinGroup", api.JoinGroups)            // 加入群聊
		v1.POST("/contact/listCommunity", api.ListCommunity)     // 群聊好友

		// 好友
		v1.POST("/search/friends", api.SearchFriends) // 好友列表
		v1.POST("/user/find", api.FindByID)           // 搜索好友
		v1.POST("/contact/addfriend", api.AddFriend)  // 添加好友

		// 聊天
		//v1.GET("/user/sendMsg", service.SendMsg)
		v1.GET("/user/sendUserMsg", service.SendUserMsg) // 发送用户消息
		v1.POST("/attach/upload", service.Upload)        // 上传文件
		v1.POST("/user/redisMsg", api.RedisMsg)          // 获取用户A、B的消息
	}
	return r
}
