package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	
	"github.com/efemirik/guardian-auth-api/internal/database"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found. Using default environment variables.")
	}

	// Initialize PostgreSQL connection pool
	log.Println("Attempting to connect to database...")
	database.Connect()

	// Initialize Gin router
	router := gin.Default()

	// Health check endpoint for infrastructure monitoring
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "Guardian Auth API is running smoothly 🛡️",
		})
	})

	// Start the server with fallback port
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" 
	}

	log.Printf("Starting server on port %s...", port)
	router.Run(":" + port)
}