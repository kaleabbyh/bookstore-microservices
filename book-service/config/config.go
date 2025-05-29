package config

import (
	"os"
)

type Config struct {
	DBHost         string
	DBUser         string
	DBPassword     string
	DBName         string
	DBPort         string
	Port           string
	JWTSecret      string
	UserServiceURL string
}

func LoadConfig() *Config {
	return &Config{
		DBHost:         getEnv("DB_HOST", "bookdb"),
		DBUser:         getEnv("DB_USER", "postgres"),
		DBPassword:     getEnv("DB_PASSWORD", "postgres"),
		DBName:         getEnv("DB_NAME", "bookdb"),
		DBPort:         getEnv("DB_PORT", "5432"),
		Port:           getEnv("PORT", "8082"),
		JWTSecret:      getEnv("JWT_SECRET", "secret"),
		UserServiceURL: getEnv("USER_SERVICE_URL", "http://user-service:8081"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}
