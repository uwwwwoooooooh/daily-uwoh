package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds all configuration for the application
type Config struct {
	DBUrl      string
	ServerPort string
	JWTSecret  string
}

// LoadConfig reads configuration from .env file or environment variables
func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	return Config{
		DBUrl:      getEnv("DATABASE_URL", "host=localhost user=postgres password=postgres dbname=dailyuwoh port=5432 sslmode=disable"),
		ServerPort: getEnv("SERVER_PORT", ":8080"),
		JWTSecret:  getEnv("JWT_SECRET", "secret"),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
