package http

import (
	"gpt/config"
	base "gpt/internal/container"
	"gpt/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, cfg *config.Config, container *base.Container) {
	api := app.Group("/api")

	auth := api.Group("/auth")

	auth.Post("/register", container.AuthHandler.Register)
	auth.Post("/login", container.AuthHandler.Login)
	auth.Post("/refresh", container.AuthHandler.Refresh)

	// Protected Routes V1
	protected := api.Group("", middleware.AuthMiddleware(container.TokenService, container.UserRepo))

	// =====================
	// User Routes
	// =====================
	userGroup := protected.Group("/users")
	userGroup.Get("/", container.UserHandler.GetAll)
	userGroup.Post("/add", container.UserHandler.Create)
	userGroup.Put("/:id", container.UserHandler.Update)
	userGroup.Delete("/:id", container.UserHandler.Delete)

	// =====================
	// Profile Routes
	// =====================
	profileGroup := protected.Group("/profile")
	profileGroup.Get("/", container.AuthHandler.Profile)

	// =====================
	// Role Routes
	// =====================
	roleGroup := protected.Group("/roles")
	roleGroup.Get("/", container.RoleHandler.GetAll)
	roleGroup.Post("/add", container.RoleHandler.Create)
	roleGroup.Put("/:id", container.RoleHandler.Update)
	roleGroup.Delete("/:id", container.RoleHandler.Delete)

	// =====================
	// Role Routes
	// =====================
	permissionGroup := protected.Group("/permission")
	permissionGroup.Get("/", container.PermissionHandler.GetAll)
	permissionGroup.Post("/add", container.PermissionHandler.Create)
	permissionGroup.Put("/:id", container.PermissionHandler.Update)
	permissionGroup.Delete("/:id", container.PermissionHandler.Delete)
}
