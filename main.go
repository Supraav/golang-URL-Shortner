package main

import (
	"treeleaf/controllers"
	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()
	router.POST("/shorten",controllers.CreateShortURL)
	router.GET("/:shortCode", controllers.GetURL)
	router.Run("localhost:8080")
}