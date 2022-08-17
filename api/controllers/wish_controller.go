package controllers

import (
	"net/http"
	"uruapi/database"
	"uruapi/models"

	"github.com/gin-gonic/gin"
)


// 出品する商品を作成
func PostWish(c *gin.Context) {

	var wish models.Wish
	if err := c.Bind(&wish); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	database.DB.Create(&wish)
}


// 編集
func EditWish(c *gin.Context) {

	var wish models.Wish
	if err := c.Bind(&wish); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	wish_id := c.Param("wish_id")
	database.DB.Model(&wish).Where("id=?", wish_id).Updates(&wish)
}


// 削除
func DeleteWish(c *gin.Context) {

	wish_id := c.Param("wish_id")
	database.DB.Delete(&models.Wish{}, wish_id)
}