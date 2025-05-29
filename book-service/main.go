package main

import (
	"github.com/kaleabbyh/book-service/config"
	"github.com/kaleabbyh/book-service/db"
	"github.com/kaleabbyh/book-service/handler"
	"github.com/labstack/echo/v4"
)

func main() {

	cfg := config.LoadConfig()

	db.Init(cfg)

	e := echo.New()

	e.POST("/books", handler.CreateBook, handler.JWTMiddleware)
	e.GET("/books", handler.GetBooks)

	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
