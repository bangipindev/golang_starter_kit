package main

import (
	"gpt/config"
	"gpt/internal/container"
	"gpt/internal/delivery/http"
	"gpt/internal/infrastructure"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	cfg := config.LoadConfig()

	db, err := infrastructure.InitDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	infrastructure.WaitForDB(db)

	dbUrl := infrastructure.BuildDBURL(cfg)

	mode := cfg.AppMode

	log.Printf("Server running in %s mode", strings.ToUpper(mode))

	if cfg.RunMigration {
		migrationPath := os.Getenv("MIGRATION_PATH")
		if migrationPath == "" {
			migrationPath = "file://migrations"
		}
		infrastructure.RunMigrations(dbUrl, migrationPath)
		infrastructure.RunSeed(db)
		log.Println("Migration & Seed executed")
	}

	container := container.NewContainer(db, cfg)

	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())

	http.SetupRoutes(app, cfg, container)

	// ======================
	// Start Server
	// ======================
	log.Println("Server started on port", cfg.AppPort)
	log.Fatal(app.Listen(":" + cfg.AppPort))
}
