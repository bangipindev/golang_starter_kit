package infrastructure

import (
	"database/sql"
	"log"

	"github.com/pressly/goose/v3"
)

func RunMigrations(db *sql.DB) {
	if err := goose.SetDialect("mysql"); err != nil {
		log.Fatal(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration executed successfully")
}
