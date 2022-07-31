package main

import (
	"log"

	//"github.com/gin-gonic/gin"
	"uruapi/database"
	"uruapi/routes"
)

func main() {
	database.Connect()
	router := routes.GetRouter()
	if err := router.Run(":8888"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
