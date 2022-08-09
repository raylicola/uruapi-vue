package controllers

import (
	"net/http"
	"uruapi/database"
	"uruapi/models"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var item models.Item
	var transaction models.Transaction
	if err := c.Bind(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	database.DB.Create(&transaction)

	database.DB.Model(&item).Where("id=?", transaction.ItemID).Update("is_sold", 1)
}