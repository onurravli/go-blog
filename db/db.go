package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/onurravli/go-blog/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	if os.Getenv("DATABASE_URL") == "" {
		panic("DATABASE_URL is not set")
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database")
	DB = db
	_err := DB.AutoMigrate(&models.Post{})
	if _err != nil {
		panic(_err)
	}
	fmt.Println("Tables created")
}
