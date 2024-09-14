package service

import (
	"IMProject/models"
	"IMProject/pkg/e"
	"IMProject/serializer"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"strconv"
)

// UserUpdateService 前端请求过来的数据
type UserUpdateService struct {
	Name     string
	Password string
	Phone    string
	Icon     string
	Email    string
}

// UserUpdate
// @Summary 修改用户
// @Tags 用户模块
// @param id formData string false "id"
// @param name formData string false "name"
// @param password formData string false "password"
// @param phone formData string false "phone"
// @param email formData string false "email"
// @Success 200 {string} json{"code","message"}
// @Router /user/updateUser [post]
func (service *UserUpdateService) UserUpdate(c *gin.Context) *serializer.Response {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = service.Name
	user.PassWord = service.Password
	user.Phone = service.Phone
	user.Avatar = service.Icon
	user.Email = service.Email

	code := e.SUCCESS

	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		code = e.ERROR_MATCHED_USERNAME
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	} else {
		models.UpdateUser(user)
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
}
