package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/efemirik/guardian-auth-api/internal/controllers"
	"github.com/efemirik/guardian-auth-api/internal/database"
	"github.com/efemirik/guardian-auth-api/internal/middleware"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found. Using default environment variables.")
	}

	// Connect to PostgreSQL and run migrations
	log.Println("Attempting to connect to database...")
	database.Connect()

	// Initialize Gin router
	router := gin.Default()

	// Public Routes (No authentication required)
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "success", "message": "API is online"})
	})

	authGroup := router.Group("/api/auth")
	{
		authGroup.POST("/register", controllers.Register)
		authGroup.POST("/login", controllers.Login)
	}

	// Protected Routes (Require valid JWT)
	protectedGroup := router.Group("/api/protected")
	protectedGroup.Use(middleware.RequireAuth) // Apply the security wall here
	{
		// Example of a route only accessible with a token
		protectedGroup.GET("/profile", func(c *gin.Context) {
			// Retrieve the userID injected by the middleware
			userID, _ := c.Get("userID")
			c.JSON(http.StatusOK, gin.H{
				"message": "Welcome to the secure zone!",
				"user_id": userID,
			})
		})
	}

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Starting server on port %s...", port)
	router.Run(":" + port)
}