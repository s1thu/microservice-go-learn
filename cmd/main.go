package main

import (
	"example/go-web-gin/config"
	"example/go-web-gin/middleware"
	"example/go-web-gin/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()

	route := gin.Default()
	// Use reusable request logging middleware
	route.Use(middleware.RequestLogger())
	// Register routes from router package
	router.RegisterRoutes(route)

	route.Run("localhost:8080")
}
