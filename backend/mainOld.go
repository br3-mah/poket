package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"pocket-ecommerce/config"
	"pocket-ecommerce/routes"
)

func mainK() {
	// Load environment variables
	config.LoadEnv()

	// Initialize database
	db := config.ConnectDB()

	// Get sql.DB object and defer closing it
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}
	defer sqlDB.Close()

	// Create Gin Router
	r := gin.Default()

	// Apply CORS and Middleware
	routes.SetupRoutes(r)

	// Start Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server running on port:", port)
	log.Fatal(r.Run(":" + port))
}
