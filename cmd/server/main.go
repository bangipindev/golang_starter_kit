package main

import (
	"gpt/config"
	"gpt/internal/delivery/http/handler"
	"gpt/internal/infrastructure"
	"gpt/internal/infrastructure/repository"
	"gpt/internal/usecase"
	"log"

	httpDelivery "gpt/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// ======================
	// Load Configuration
	// ======================
	cfg := config.LoadConfig()

	// ======================
	// Initialize Database
	// ======================
	db, err := infrastructure.InitDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// ✅ Only run migration in dev
	if cfg.AppMode == "dev" {
		infrastructure.RunMigrations(db)
		log.Println("Congrats...! Server Running in DEV mode")
	} else if cfg.AppMode == "staging" {
		infrastructure.RunMigrations(db)
		log.Println("Congrats...! Server Running in STAGING mode")
	} else {
		log.Println("Congrats...! Server Running in PRODUCTION mode")
	}

	// ======================
	// Dependency Injection
	// ======================
	userRepo := repository.NewUserRepository(db)
	authUsecase := usecase.NewAuthUsecase(userRepo, cfg.JWTSecret, cfg.JWTAccessExpiry, cfg.JWTRefreshExpiry)
	authHandler := handler.NewAuthHandler(authUsecase)

	// ======================
	// Initialize Fiber
	// ======================
	app := fiber.New()

	// Global Middlewares
	app.Use(logger.New())
	app.Use(recover.New())

	// ======================
	// Setup Routes
	// ======================
	httpDelivery.SetupRoutes(app, authHandler)

	// ======================
	// Start Server
	// ======================
	log.Fatal(app.Listen(":" + cfg.AppPort))
}
