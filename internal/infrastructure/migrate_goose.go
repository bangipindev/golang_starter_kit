package infrastructure

import (
	"database/sql"
	"log"

	"github.com/pressly/goose/v3"
)

func RunMigrationsGoose(db *sql.DB) {
	if err := goose.SetDialect("mysql"); err != nil {
		log.Fatal(err)
	}

	wd, _ := os.Getwd()
	migrationPath := wd + "/migrations"

	if err := goose.Up(db, migrationPath); err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration executed successfully")
}
