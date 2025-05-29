package main

import (
	"github.com/kaleabbyh/user-service/config"
	"github.com/kaleabbyh/user-service/db"
	"github.com/kaleabbyh/user-service/handler"
	"github.com/labstack/echo/v4"
)

func main() {

	cfg := config.LoadConfig()

	db.Init(cfg)

	e := echo.New()

	e.POST("/register", handler.Register)
	e.POST("/login", handler.Login)
	e.GET("/users/:id", handler.GetUser)

	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
