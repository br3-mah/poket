package models

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDB initializes the database connection and runs migrations
func ConnectDB() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		dsn = "user:password@tcp(127.0.0.1:3306)/pcommerce?charset=utf8mb4&parseTime=True&loc=Local" // Example DSN
	}

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to the database:", err)
	}
	fmt.Println("Database connected successfully")

	// Run the migrations (create tables, indexes, etc.)
	err = DB.AutoMigrate(&User{}) // Automatically migrate the User model
	if err != nil {
		log.Fatal("failed to migrate the database:", err)
	}
	fmt.Println("Database migrated successfully")
}
