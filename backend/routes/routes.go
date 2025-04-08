package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"pocket-ecommerce/controllers/auth"
	"time"
)

func SetupRoutes(router *gin.Engine) {
	// Single CORS configuration
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"*"}, // Simplified for debugging
		AllowHeaders:     []string{"*"}, // Simplified for debugging
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// API routes with absolute path
	router.POST("/api/register", controllers.RegisterUser)
	router.POST("/api/login", controllers.LoginUser)

	// OPTIONAL: Add test route for debugging
	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
}