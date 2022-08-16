package routes

import (
	"time"
	"uruapi/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		// 許可するアクセス元
		AllowOrigins: []string{
			"http://localhost:8080",
		},
		// アクセス許可するHTTPメソッド
		AllowMethods: []string{
			"POST",
			"GET",
			"PUT",
			"DELETE",
			"HEAD",
		},
		// 許可するHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Content-Type",
			"Accept",
		},
		// cookieなどの情報を必要とする
		AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))

	// Item
	router.POST("/item/create", controllers.CreateItem)
	router.PUT("/item/edit/:item_id", controllers.EditItem)
	router.DELETE("/item/delete/:item_id", controllers.DeleteItem)

	// Wish
	router.POST("/wish/create", controllers.CreateWish)
	router.PUT("/wish/edit/:wish_id", controllers.EditWish)
	router.DELETE("/wish/delete/:wish_id", controllers.DeleteWish)

	// Chat
	router.POST("/chat/create", controllers.CreateChat)

	// Review
	router.POST("/review/create", controllers.CreateReview)

	// Sale
	router.POST("/sale/create", controllers.CreateSale)

	// Transaction
	router.POST("/transaction/create", controllers.CreateTransaction)

	// Get Data
	router.GET("/seller/:seller_id", controllers.GetSellerItem)
	router.GET("/seller/review/:seller_id", controllers.GetSellerReview)
	router.GET("/wish/:wish_id", controllers.GetWish)
	router.GET("/item/:item_id", controllers.GetItem)
	router.GET("/wish", controllers.GetAllWish)
	router.GET("/wish/search", controllers.GetSearchedWish)
	router.GET("/user/:user_id/wish", controllers.GetMyWish)
	router.GET("/user/:user_id/item", controllers.GetMyItem)
	router.GET("/user/:user_id/purchased", controllers.GetMyPurchasedItem)
	router.GET("/user/:user_id/sold", controllers.GetMySoldItem)
	return router
}
