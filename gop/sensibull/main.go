package main

import (
	"Go/src/go/go-sensibull/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Redis client
	//utils.InitRedis("localhost:6379", "", 0)

	// Create a Gin router
	router := gin.Default()

	// Define the GET API route and its handler
	router.GET("/get-data", handlers.GetHandler)

	// Define the POST API route and its handler
	router.POST("/post-data", handlers.PostHandler)

	// Start the server
	router.Run(":8080")
}
