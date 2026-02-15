package http

import (
	"gpt/config"
	"gpt/internal/container"
	"gpt/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, cfg *config.Config, container *container.Container) {
	api := app.Group("/api")

	// =========================
	// API VERSION 1
	// =========================
	v1 := api.Group("/v1")
	/*
	* NO AUTH
	 */
	authV1 := v1.Group("/auth")
	authV1.Post("/register", container.AuthHandler.Register)
	authV1.Post("/login", container.AuthHandler.Login)
	authV1.Post("/refresh", container.AuthHandler.Refresh)

	// Protected Routes V1
	protectedV1 := v1.Group("", middleware.AuthMiddleware(container.TokenService))

	// =====================
	// Profile Routes
	// =====================
	profileGroup := protectedV1.Group("/profile")
	profileGroup.Get("/", container.AuthHandler.Profile)

	// =====================
	// Role Routes
	// =====================
	roleGroup := protectedV1.Group("/roles")
	roleGroup.Get("/", container.RoleHandler.GetRoles)
	roleGroup.Get("/add", container.RoleHandler.Add)
}
