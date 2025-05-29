package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kaleabbyh/book-service/db"
	"github.com/kaleabbyh/book-service/model"
	"github.com/labstack/echo/v4"
)

var jwtSecret = []byte("secret")

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.ErrUnauthorized
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil {
			return echo.ErrUnauthorized

		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user_id", claims["user_id"])
			return next(c)
		}

		return echo.ErrUnauthorized
	}
}

func userExists(userID string) bool {
	resp, err := http.Get("http://user-service:8081/users/" + userID)
	if err != nil || resp.StatusCode != http.StatusOK {
		return false
	}
	defer resp.Body.Close()
	var result map[string]interface{}
	return json.NewDecoder(resp.Body).Decode(&result) == nil
}

func CreateBook(c echo.Context) error {
	book := new(model.Book)
	if err := c.Bind(book); err != nil {
		return err
	}

	userID := fmt.Sprintf("%.0f", c.Get("user_id").(float64))
	if !userExists(userID) {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid user"})
	}

	if err := db.DB.Create(book).Error; err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, book)
}

func GetBooks(c echo.Context) error {
	var books []model.Book
	db.DB.Find(&books)
	return c.JSON(http.StatusOK, books)
}
