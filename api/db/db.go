package db

import (
	"fmt"
	"log"
	"os"

	"github.com/44t4nk1/StudentPortal/api/models"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var (
	DB *gorm.DB
)

func GetDB() *gorm.DB {
	return DB
}

func RunDB() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	}

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PASSWORD"))
	DB, err = gorm.Open(os.Getenv("DB_DRIVER"), DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", os.Getenv("DB_DRIVER"))
		log.Fatal("This is the error:", err)
	}

	DB.Debug().AutoMigrate(&models.Student{})

	return DB, err
}
