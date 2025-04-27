package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set Gin mode based on environment variable (optional)
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()

	// Simple health check / ping endpoint
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Add other routes here later...
	// e.g., API group
	// api := router.Group("/api/v1")
	// {
		// Auth routes
		// Portfolio routes
		// Upload routes
	// }


	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not set
	}
	if err := router.Run(":" + port); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
