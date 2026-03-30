package main

import (
	"log"
	"os"
	
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	
	// DİKKAT: Eğer go.mod dosyanın en üstünde "module github.com/.../guardian-auth-api" yazıyorsa, burayı ona göre güncellemelisin.
	// Eğer sadece "module guardian-auth-api" yazıyorsa, bu satır tam olarak böyle kalmalı:
	
	"github.com/efemirik/guardian-auth-api/internal/database"
)

func main() {
	// 1. .env dosyasını yükle
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found. Using default environment variables.")
	}

	// 2. Veritabanı bağlantısını başlat (internal/database/db.go içindeki Connect fonksiyonunu çağırıyoruz)
	log.Println("Attempting to connect to database...")
	database.Connect()

	// 3. Gin Router'ı oluştur
	router := gin.Default()

	// 4. Health Check (Sağlık Kontrolü) Endpoint'i
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "Guardian Auth API is running smoothly 🛡️",
		})
	})

	// 5. Sunucuyu Başlat
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Fallback port
	}

	log.Printf("Starting server on port %s...", port)
	router.Run(":" + port)
}