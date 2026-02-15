package main

import (
	"gpt/config"
	"gpt/internal/container"
	"gpt/internal/delivery/http"
	"gpt/internal/infrastructure"
	"log"

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

	if cfg.AppMode == "dev" {
		infrastructure.RunMigrations(db)
		log.Println("Congrats...! Server Running in DEV mode")
	} else if cfg.AppMode == "staging" {
		infrastructure.RunMigrations(db)
		log.Println("Congrats...! Server Running in STAGING mode")
	} else {
		log.Println("Congrats...! Server Running in PRODUCTION mode")
	}

	container := container.NewContainer(db, cfg)

	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())

	http.SetupRoutes(app, cfg, container)

	// ======================
	// Start Server
	// ======================
	log.Fatal(app.Listen(":" + cfg.AppPort))
}
