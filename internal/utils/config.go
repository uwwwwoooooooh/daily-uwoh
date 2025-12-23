package utils

import (
	"log"

	"github.com/spf13/viper"
)

// Config holds all configuration for the application
type Config struct {
	DBUrl              string `mapstructure:"DATABASE_URL"`
	ServerPort         string `mapstructure:"SERVER_PORT"`
	JWTSecret          string `mapstructure:"JWT_SECRET"`
	JWTExpirationHours int    `mapstructure:"JWT_EXPIRATION_HOURS"`
}

// LoadConfig reads configuration from .env file or environment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return
		}
		log.Println("No .env file found, using system environment variables")
	}

	err = viper.Unmarshal(&config)
	return
}
