package service

import (
	"IMProject/models"
	"strconv"
	"text/template"

	"github.com/gin-gonic/gin"
)

// GetIndex
// @Tags 首页
// @Success 200 {string} welcome
// @Router /index [get]
func GetIndex(c *gin.Context) {
	// template.ParseFiles 是 Go 的 html/template 包中的一个函数，用于解析一个或多个 HTML 模板文件。
	// 它将 index.html 和 views/chat/head.html 文件解析为模板。
	ind, err := template.ParseFiles("index.html", "views/chat/head.html")
	if err != nil {
		panic(err)
	}
	// 用于将解析后的模板渲染到 HTTP 响应中。
	// "index" 是传递给模板的数据，在模板中可以使用该数据进行动态内容渲染。
	// 这里传递的数据是一个字符串 "index"，可以在 index.html 模板中引用。
	err = ind.Execute(c.Writer, "index")
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{
		"message": "welcome !!  ",
	})
}

func ToRegister(c *gin.Context) {
	ind, err := template.ParseFiles("views/user/register.html")
	if err != nil {
		panic(err)
	}
	err = ind.Execute(c.Writer, "register")
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{
		"message": "welcome !!  ",
	})
}

func ToChat(c *gin.Context) {
	ind, err := template.ParseFiles("views/chat/index.html",
		"views/chat/head.html",
		"views/chat/foot.html",
		"views/chat/tabmenu.html",
		"views/chat/concat.html",
		"views/chat/group.html",
		"views/chat/profile.html",
		"views/chat/createcom.html",
		"views/chat/userinfo.html",
		"views/chat/main.html")
	if err != nil {
		panic(err)
	}
	userId, _ := strconv.Atoi(c.Query("userId"))
	token := c.Query("token")
	user := models.UserBasic{}
	user.ID = uint(userId)
	user.Identity = token
	ind.Execute(c.Writer, user)
	c.JSON(200, gin.H{
		"message": "welcome !!  ",
	})
}
