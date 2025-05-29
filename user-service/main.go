package main

import (
	"github.com/kaleabbyh/user-service/db"
	"github.com/kaleabbyh/user-service/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	db.Init()
	e := echo.New()

	e.POST("/register", handler.Register)
	e.POST("/login", handler.Login)
	e.GET("/users/:id", handler.GetUser)

	e.Logger.Fatal(e.Start(":8081"))
}
