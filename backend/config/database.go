package config

import (
	"fmt"
	"log"
	"os"

	
	"pocket-ecommerce/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found")
	} else {
		log.Println("Found Env")
	}
}

func ConnectDB() *gorm.DB {
	// Correct DSN format for MySQL
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}
	fmt.Println("✅ Database connected successfully!")

	

	// Run the migrations (create tables, indexes, etc.)
	err = DB.AutoMigrate(&models.User{})  // Automatically migrate the User model
	if err != nil {
		log.Fatal("failed to migrate the database:", err)
	}
	fmt.Println("Database migrated successfully")
	return DB
}
