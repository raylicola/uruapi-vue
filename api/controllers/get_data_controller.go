package controllers

import (
	"net/http"
	"uruapi/database"
	"uruapi/models"

	"github.com/gin-gonic/gin"
)

// Wishを1件取得
func GetWish(c *gin.Context) {
	wish_id := c.Param("wish_id")
	var wish models.Wish
	var items []models.Item

	database.DB.Where("id = ?", wish_id).First(&wish)
	database.DB.Joins("left join sales on sales.item_id = items.id").Where("wish_id = ?", wish_id).Find(&items)
	c.JSON(http.StatusOK, gin.H{"wish": wish, "items": items})
}

// Item(+Chat, Review)を1件取得
func GetItem(c *gin.Context) {
	item_id := c.Param("item_id")
	var item models.Item
	var chats []models.Chat
	var reviews []models.Review

	database.DB.Where("id = ?", item_id).First(&item)
	database.DB.Where("reviewee_id=?", item.UserID).Find(&reviews)
	database.DB.Where("item_id=?", item_id).Find(&chats)

	c.JSON(http.StatusOK, gin.H{"item": item, "chats": chats, "reviews": reviews})
}

// Wishを全件取得
func GetAllWish(c *gin.Context) {
	var wishes []models.Wish
	database.DB.Find(&wishes)
	c.JSON(http.StatusOK, gin.H{"wishes": wishes})
}


// Wishをキーワード検索
func GetSearchedWish(c *gin.Context) {
	keyword := "%"+c.Query("keyword")+"%"
	var wishes []models.Wish
	database.DB.Where("title LIKE ?", keyword).Or("detail LIKE ?", keyword).Find(&wishes)
	c.JSON(http.StatusOK, gin.H{"wishes": wishes})
}


// 自分のWishを取得
func GetMyWish(c *gin.Context) {
	user_id := c.Param("user_id")
	var wishes []models.Wish
	database.DB.Where("user_id = ?", user_id).Find(&wishes)
	c.JSON(http.StatusOK, gin.H{"wishes": wishes})
}

// 自分のItemを取得
func GetMyItem(c *gin.Context) {
	user_id := c.Param("user_id")
	var items []models.Item
	database.DB.Where("user_id = ?", user_id).Find(&items)
	c.JSON(http.StatusOK, gin.H{"items": items})
}

// 購入履歴を取得
func GetMyPurchasedItem(c *gin.Context) {
	var purchasedItems []models.Item
	user_id := c.Param("user_id")
	database.DB.Raw("SELECT * FROM items INNER JOIN transactions ON items.id = transactions.item_id WHERE 	purchaser_id = ?", user_id).Scan(&purchasedItems)

	c.JSON(http.StatusOK, gin.H{"purchasedItems": purchasedItems})
}

// 販売履歴を取得
func GetMySoldItem(c *gin.Context) {
	var soldItems []models.Item
	user_id := c.Param("user_id")
	database.DB.Raw("SELECT * FROM items INNER JOIN transactions ON items.id = transactions.item_id WHERE items.user_id = ?", user_id).Scan(&soldItems)

	c.JSON(http.StatusOK, gin.H{"soldItems": soldItems})
}