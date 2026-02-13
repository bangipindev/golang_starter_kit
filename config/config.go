package config

import "github.com/spf13/viper"

type Config struct {
	AppMode   string
	DBHost    string
	DBPort    string
	DBUser    string
	DBPass    string
	DBName    string
	JWTSecret string
	AppPort   string
}

func LoadConfig() *Config {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	return &Config{
		AppMode:   viper.GetString("APP_MODE"),
		DBHost:    viper.GetString("DB_HOST"),
		DBPort:    viper.GetString("DB_PORT"),
		DBUser:    viper.GetString("DB_USER"),
		DBPass:    viper.GetString("DB_PASS"),
		DBName:    viper.GetString("DB_NAME"),
		JWTSecret: viper.GetString("JWT_SECRET"),
		AppPort:   viper.GetString("APP_PORT"),
	}
}
