package database

import (
	"backend/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase initializes the database connection
// @Description Establishes connection to PostgreSQL database using environment variables
func ConnectDatabase() {
	// Try to load .env from different possible locations
	envPaths := []string{
		".env",      // Current directory
		"../.env",   // Parent directory
		"/app/.env", // Docker volume mount point
	}

	envLoaded := false
	for _, path := range envPaths {
		if err := godotenv.Load(path); err == nil {
			log.Printf("Loaded .env from %s", path)
			envLoaded = true
			break
		}
	}

	if !envLoaded {
		log.Printf("Warning: .env file not found, using environment variables")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Run migrations directly here instead of importing migrations package
	log.Println("Running database migrations...")
	if err := db.AutoMigrate(&models.User{}, &models.Book{}); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}
	log.Println("Database migrations completed successfully")

	DB = db
	log.Println("Database connected and migrations completed successfully")
}
