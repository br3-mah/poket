package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"pocket-ecommerce/config"
	"pocket-ecommerce/routes"
	
	"pocket-ecommerce/middleware" // Add this import
)

func main() {
	config.LoadEnv()
	db := config.ConnectDB()

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}else{
		fmt.Println("Database connected...")
	}
	defer sqlDB.Close()

	r := gin.Default()
	r.Use(middleware.Logger())
	// REMOVED CORS FROM HERE - Let routes.go handle it
	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Server running on port:", port)
	log.Fatal(r.Run(":" + port))
}