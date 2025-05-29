package db

import (
	"log"

	"github.com/kaleabbyh/user-service/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dsn := "host=userdb user=postgres password=postgres dbname=userdb port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to user db: ", err)
	}

	DB.AutoMigrate(&model.User{})
}
