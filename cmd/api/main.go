package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 1. .env dosyasını yükle
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found. Using default environment variables.")
	}

	// 2. Gin Router'ı oluştur
	router := gin.Default()

	// 3. Health Check (Sağlık Kontrolü) Endpoint'i - DevOps uzmanları buna bayılır!
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "Guardian Auth API is running smoothly 🛡️",
		})
	})

	// 4. Sunucuyu Başlat
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Fallback port
	}

	log.Printf("Starting server on port %s...", port)
	router.Run(":" + port)
}