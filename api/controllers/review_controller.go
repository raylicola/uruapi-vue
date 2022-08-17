package controllers

import (
	"net/http"
	"uruapi/database"
	"uruapi/models"

	"github.com/gin-gonic/gin"
)

// レビューを作成
func PostReview(c *gin.Context) {
	var review models.Review
	if err := c.Bind(&review); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	database.DB.Create(&review)
}