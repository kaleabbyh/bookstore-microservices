package db

import (
	"fmt"
	"log"

	"github.com/kaleabbyh/book-service/config"
	"github.com/kaleabbyh/book-service/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(cfg *config.Config) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to user db: ", err)
	}

	if err := DB.AutoMigrate(&model.Book{}); err != nil {
		log.Fatal("failed to migrate book model: ", err)
	}
}
