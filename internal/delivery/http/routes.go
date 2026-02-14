package http

import (
	"gpt/config"
	"gpt/internal/delivery/http/handler"
	"gpt/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, authHandler *handler.AuthHandler) {
	cfg := config.LoadConfig()

	// =========================
	// Base API
	// =========================
	api := app.Group("/api")

	// =========================
	// API VERSION 1
	// =========================
	v1 := api.Group("/v1")

	// Auth Routes V1
	authV1 := v1.Group("/auth")
	authV1.Post("/register", authHandler.Register)
	authV1.Post("/login", authHandler.Login)
	authV1.Post("/refresh", authHandler.Refresh)

	// Protected Routes V1
	protectedV1 := v1.Group("", middleware.JWTProtected(cfg.JWTSecret))
	protectedV1.Get("/profile", authHandler.Profile)

	// ======================================================
	// API VERSION 2 (Contoh Future Development)
	// ======================================================

	/*
		v2 := api.Group("/v2")

		// Misalnya nanti kita ubah struktur response
		authV2 := v2.Group("/auth")
		authV2.Post("/register", authHandler.RegisterV2)
		authV2.Post("/login", authHandler.LoginV2)

		// Protected Routes V2
		protectedV2 := v2.Group("", middleware.JWTProtected())
		protectedV2.Get("/profile", func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"status":  "success",
				"message": "JWT valid 🔐 (v2)",
				"version": "2.0",
			})
		})

		// Contoh perubahan di v2:
		// - Response format berbeda
		// - Tambah refresh token
		// - Role-based access control
		// - Pagination standard baru
	*/
}
