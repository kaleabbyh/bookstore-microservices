package db

import (
	"log"

	"github.com/kaleabbyh/book-service/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := "host=bookdb user=postgres password=postgres dbname=bookdb port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to book db: ", err)
	}

	DB.AutoMigrate(&model.Book{})
}
