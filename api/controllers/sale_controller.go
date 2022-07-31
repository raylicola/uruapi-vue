package controllers

import (
	"net/http"
	"uruapi/database"
	"uruapi/models"

	"github.com/gin-gonic/gin"
)

// レビューを作成
func CreateSale(c *gin.Context) {
	var sale models.Sale
	if err := c.Bind(&sale); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	database.DB.Create(&sale)
}