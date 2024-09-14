package service

import (
	"IMProject/models"
	"github.com/gin-gonic/gin"
)

func SendUserMsg(c *gin.Context) {
	models.Chat(c.Writer, c.Request)
}
