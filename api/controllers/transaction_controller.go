package controllers

import (
	"net/http"
	"uruapi/database"
	"uruapi/models"

	"github.com/gin-gonic/gin"
)

func PostTransaction(c *gin.Context) {
	var item models.Item
	var transaction models.Transaction
	if err := c.Bind(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	database.DB.Model(&item).Where("id=?", transaction.ItemID).Update("purchaser_id", transaction.PurchaserID)
}