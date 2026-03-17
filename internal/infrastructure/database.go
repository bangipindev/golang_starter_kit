package infrastructure

import (
	"database/sql"
	"fmt"
	"gpt/config"
	"log"
	"time"

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

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// APPLY CONNECTION POOL
	db.SetMaxIdleConns(cfg.DBPool.MaxIdleConnections)
	db.SetMaxOpenConns(cfg.DBPool.MaxOpenConnections)
	db.SetConnMaxIdleTime(cfg.DBPool.MaxIdleTime)
	db.SetConnMaxLifetime(cfg.DBPool.MaxConnectionLifetime)

	return db, nil
}

func WaitForDB(db *sql.DB) {
	for i := 0; i < 10; i++ {
		if err := db.Ping(); err == nil {
			log.Println("Database connected")
			return
		}
		log.Println("Waiting for DB...")
		time.Sleep(2 * time.Second)
	}
	log.Fatal("Database not ready")
}
