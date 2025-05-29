package main

import (
	"github.com/kaleabbyh/book-service/db"
	"github.com/kaleabbyh/book-service/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	db.Init()
	e := echo.New()

	e.POST("/books", handler.CreateBook, handler.JWTMiddleware)
	e.GET("/books", handler.GetBooks)

	e.Logger.Fatal(e.Start(":8082"))
}
