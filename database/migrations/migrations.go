package main

import (
	"backend/models"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// RunMigrations executes database migrations
func RunMigrations(db *gorm.DB) error {
	log.Println("Starting database migrations...")
	if err := db.AutoMigrate(&models.User{}, &models.Book{}); err != nil {
		log.Printf("Failed to run migrations: %v", err)
		return err
	}
	log.Println("Migrations completed successfully")
	return nil
}

func main() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"))

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := RunMigrations(db); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}
