package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB, projenin her yerinden veritabanına erişmek için kullanacağımız global değişken
var DB *gorm.DB

func Connect() {
	// .env dosyasından bilgileri çekiyoruz
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// PostgreSQL bağlantı stringini (DSN) oluşturuyoruz
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", 
		host, user, password, dbname, port)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
	}

	log.Println("Successfully connected to the PostgreSQL database! 🐘")
}