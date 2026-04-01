package database

import (
	"fmt"
	"log"
	"os"

	"github.com/efemirik/guardian-auth-api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB represents the global database connection pool
var DB *gorm.DB

// Connect initializes the PostgreSQL connection and runs schemas migrations
func Connect() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// Construct the Data Source Name (DSN)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", 
		host, user, password, dbname, port)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
	}

	log.Println("Successfully connected to the PostgreSQL database! 🐘")

	// Auto-migrate database schemas
	err = DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
	
	log.Println("Database migration completed! Tables are ready.")
}