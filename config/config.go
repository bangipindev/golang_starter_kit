package config

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	AppMode string
	AppPort string

	RunMigration     bool
	DBHost           string
	DBPort           string
	DBUser           string
	DBPass           string
	DBName           string
	JWTSecret        string
	JWTAccessExpiry  time.Duration
	JWTRefreshExpiry time.Duration

	DBPool DBConnectionPoolConfig
}

type DBConnectionPoolConfig struct {
	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxIdleTime           time.Duration
	MaxConnectionLifetime time.Duration
}

func LoadConfig() *Config {
	_ = godotenv.Load()
	viper.AutomaticEnv()
	accessExp, _ := time.ParseDuration(viper.GetString("JWT_ACCESS_EXPIRY"))
	refreshExp, _ := time.ParseDuration(viper.GetString("JWT_REFRESH_EXPIRY"))

	if accessExp == 0 {
		log.Fatal("JWT_ACCESS_EXPIRY not set")
	}

	if refreshExp == 0 {
		log.Fatal("JWT_REFRESH_EXPIRY not set")
	}

	return &Config{
		AppMode: viper.GetString("APP_MODE"),
		AppPort: viper.GetString("APP_PORT"),

		RunMigration:     viper.GetBool("RUN_MIGRATION"),
		DBHost:           viper.GetString("DB_HOST"),
		DBPort:           viper.GetString("DB_PORT"),
		DBUser:           viper.GetString("DB_USER"),
		DBPass:           viper.GetString("DB_PASS"),
		DBName:           viper.GetString("DB_NAME"),
		JWTSecret:        viper.GetString("JWT_SECRET"),
		JWTAccessExpiry:  accessExp,
		JWTRefreshExpiry: refreshExp,

		DBPool: DBConnectionPoolConfig{
			MaxIdleConnections:    viper.GetInt("MAX_IDLE_CONNECTIONS"),
			MaxOpenConnections:    viper.GetInt("MAX_OPEN_CONNECTIONS"),
			MaxIdleTime:           time.Duration(viper.GetInt("MAX_IDLE_TIME")) * time.Minute,
			MaxConnectionLifetime: time.Duration(viper.GetInt("MAX_CONNECTION_LIFETIME")) * time.Minute,
		},
	}
}
