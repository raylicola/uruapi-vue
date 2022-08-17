package controllers

import (
	"net/http"
	"uruapi/database"
	"uruapi/models"
	"github.com/gin-gonic/gin"
)

// お気に入り登録
func PostFavorite(c *gin.Context) {
	var favorite models.Favorite
	var item models.Item
	if err := c.Bind(&favorite); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	database.DB.Where("id = ?", favorite.ItemID).First(&item)
	favorite.Title = item.Title
	favorite.Price = item.Price
	favorite.Img = item.Img
	database.DB.Create(&favorite)
}

func DeleteFavorite(c *gin.Context) {
	favorite_id := c.Param("favorite_id")
	database.DB.Delete(&models.Favorite{}, favorite_id)
}