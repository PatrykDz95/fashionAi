package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var database *gorm.DB

//func init() {
//	err := godotenv.Load()
//	if err != nil {
//		log.Fatal("Error loading .env file")
//	}
//}

func InitDB() *gorm.DB {
	host := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=9920 sslmode=disable TimeZone=Asia/Shanghai",
		host, dbUser, dbPassword, dbName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	err = db.AutoMigrate()

	if err != nil {
		log.Fatal("Failed to auto-migrate the database:", err)
	}
	database = db
	return db
}

type Services struct {
	//CompanyService   *company.Service
	//ProductService   *product.Service
	//InventoryService *inventory.Service
}

func InitServices(db *gorm.DB) *Services {
	return &Services{
		//CompanyService:   company.NewCompanyService(db),
		//ProductService:   product.NewProductService(db),
		//InventoryService: inventory.NewInventoryService(db),
	}
}
