package handler

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kaleabbyh/user-service/db"
	"github.com/kaleabbyh/user-service/model"
	"github.com/labstack/echo/v4"
)

var jwtSecret = []byte("secret")

func Register(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}
	if err := db.DB.Create(user).Error; err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, user)
}

func Login(c echo.Context) error {
	creds := new(model.User)
	if err := c.Bind(creds); err != nil {
		return err
	}

	var user model.User
	if err := db.DB.Where("email = ? AND password = ?", creds.Email, creds.Password).First(&user).Error; err != nil {
		return echo.ErrUnauthorized
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenStr, err := token.SignedString(jwtSecret)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": tokenStr,
	})
}

func GetUser(c echo.Context) error {
	id := c.Param("id")
	var user model.User
	if err := db.DB.First(&user, id).Error; err != nil {
		return echo.ErrNotFound
	}
	return c.JSON(http.StatusOK, user)
}
