package infrastructure

import (
	"database/sql"
	"fmt"
	"gpt/config"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	return sql.Open("mysql", dsn)
}
