package controllers

import (
	"net/http"
	"uruapi/database"
	"uruapi/models"

	"github.com/gin-gonic/gin"
)

// チャットを作成
func CreateChat(c *gin.Context) {
	var chat models.Chat
	if err := c.Bind(&chat); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	database.DB.Create(&chat)
}