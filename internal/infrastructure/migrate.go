package infrastructure

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(dbURL, action string) {
	if action == "" {
		action = "up"
	}

	migrationPath := os.Getenv("MIGRATION_PATH")
	if migrationPath == "" {
		migrationPath = "file:///app/migrations"
	}

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

	switch action {
	case "up":
		err = m.Up()

	case "down":
		err = m.Steps(-1)

	case "reset":
		err = m.Down()

	case "version":
		version, dirty, err := m.Version()
		if err == migrate.ErrNilVersion {
			log.Println("No migration applied yet")
			return
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Current version: %d | dirty: %v", version, dirty)
		return

	default:
		log.Fatalf("Invalid migration action: %s", action)
	}

	if err != nil && err != migrate.ErrNoChange {
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
