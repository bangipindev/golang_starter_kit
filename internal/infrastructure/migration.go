package infrastructure

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(dbURL, migrationPath string) {
	m, err := migrate.New(migrationPath, dbURL)
	if err != nil {
		log.Fatal("Migration init failed:", err)
	}
	defer func() {
		srcErr, dbErr := m.Close()
		if srcErr != nil || dbErr != nil {
			log.Println("Migration close error:", srcErr, dbErr)
		}
	}()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Migration failed:", err)
	}

	version, dirty, err := m.Version()
	if err == migrate.ErrNilVersion {
		log.Println("No migration applied yet")
	} else if err != nil {
		log.Println("Version check failed:", err)
	} else {
		log.Printf("Migration success | version: %d | dirty: %v", version, dirty)
	}
}
