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

// Item(+Chat)を1件取得
func GetItem(c *gin.Context) {
	item_id := c.Param("item_id")
	var item models.Item
	var chats []models.Chat

	database.DB.Where("id = ?", item_id).First(&item)
	database.DB.Where("item_id=?", item_id).Find(&chats)

	c.JSON(http.StatusOK, gin.H{"item": item, "chats": chats})
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

// 該当ユーザーのItemを取得
func GetSellerItem(c *gin.Context) {
	seller_id := c.Param("seller_id")
	var items []models.Item
	database.DB.Where("seller_id = ?", seller_id).Find(&items)
	c.JSON(http.StatusOK, gin.H{"items": items})
}

// 該当ユーザーのReviewを取得
func GetSellerReview(c *gin.Context) {
	var reviews []models.Review
	seller_id := c.Param("seller_id")
	database.DB.Where("reviewee_id=?", seller_id).Find(&reviews)
	c.JSON(http.StatusOK, gin.H{"reviews": reviews})
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
	database.DB.Where("seller_id = ?", user_id).Find(&items)
	c.JSON(http.StatusOK, gin.H{"items": items})
}

// 購入履歴を取得
func GetMyPurchasedItem(c *gin.Context) {
	var Items []models.Item
	user_id := c.Param("user_id")
	database.DB.Where("purchaser_id=?", user_id).Find(&Items)

	c.JSON(http.StatusOK, gin.H{"Items": Items})
}

// 販売履歴を取得
func GetMySoldItem(c *gin.Context) {
	var Items []models.Item
	user_id := c.Param("user_id")
	database.DB.Where("seller_id=?", user_id).Not("purchaser_id=?","").Find(&Items)

	c.JSON(http.StatusOK, gin.H{"Items": Items})
}