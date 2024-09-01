package db

import (
	"fasion.ai/server/auth"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func InitDB() *gorm.DB {
	host := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if host == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		log.Fatal("Database connection details are not set in environment variables")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=9920 sslmode=disable TimeZone=Asia/Shanghai",
		host, dbUser, dbPassword, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	err = db.AutoMigrate(&auth.User{})

	if err != nil {
		log.Fatal("Failed to auto-migrate the database:", err)
	}
	return db
}

type Services struct {
	UserService *auth.Service
}

func InitServices(db *gorm.DB) *Services {
	return &Services{
		UserService: auth.NewUserService(db),
	}
}
