package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	AppMode          string
	DBHost           string
	DBPort           string
	DBUser           string
	DBPass           string
	DBName           string
	JWTSecret        string
	AppPort          string
	JWTAccessExpiry  time.Duration
	JWTRefreshExpiry time.Duration
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	accessExp, _ := time.ParseDuration(viper.GetString("JWT_ACCESS_EXPIRY"))
	refreshExp, _ := time.ParseDuration(viper.GetString("JWT_REFRESH_EXPIRY"))

	if accessExp == 0 {
		log.Fatal("JWT_ACCESS_EXPIRY not set")
	}

	if refreshExp == 0 {
		log.Fatal("JWT_REFRESH_EXPIRY not set")
	}

	return &Config{
		AppMode:          viper.GetString("APP_MODE"),
		DBHost:           viper.GetString("DB_HOST"),
		DBPort:           viper.GetString("DB_PORT"),
		DBUser:           viper.GetString("DB_USER"),
		DBPass:           viper.GetString("DB_PASS"),
		DBName:           viper.GetString("DB_NAME"),
		JWTSecret:        viper.GetString("JWT_SECRET"),
		AppPort:          viper.GetString("APP_PORT"),
		JWTAccessExpiry:  accessExp,
		JWTRefreshExpiry: refreshExp,
	}
}
