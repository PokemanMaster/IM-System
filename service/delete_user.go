package service

import (
	"IMProject/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

// DeleteUser
// @Summary 删除用户
// @Tags 用户模块
// @param id query string false "id"
// @Success 200 {string} json{"code","message"}
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(user)
	c.JSON(200, gin.H{
		"code":    0, //  0成功   -1失败
		"message": "删除用户成功！",
		"data":    user,
	})
}
