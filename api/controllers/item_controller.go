package controllers

import (
	"net/http"
	"uruapi/database"
	"uruapi/models"

	"github.com/gin-gonic/gin"
)


// 出品する商品を作成
func CreateItem(c *gin.Context) {

	var item models.Item
	if err := c.Bind(&item); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	database.DB.Create(&item)
}


// 編集
func EditItem(c *gin.Context) {

	var item models.Item
	if err := c.Bind(&item); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	item_id := c.Param("item_id")
	database.DB.Model(&item).Where("id=?", item_id).Updates(&item)
}


// 削除
func DeleteItem(c *gin.Context) {

	item_id := c.Param("item_id")
	database.DB.Delete(&models.Item{}, item_id)
}