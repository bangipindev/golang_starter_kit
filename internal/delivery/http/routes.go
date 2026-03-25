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
	userGroup := protected.Group("/users", middleware.RequireRole("admin"))
	userGroup.Get("/", container.UserHandler.GetAll)
	userGroup.Post("/add", container.UserHandler.Create)
	userGroup.Post("/:id/roles", container.UserHandler.AssignRole)
	userGroup.Post("/:id/permissions", container.UserHandler.AssignPermission)
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
	roleGroup.Get("/", middleware.RequirePermission("view_roles"), container.RoleHandler.GetAll)
	roleGroup.Post("/add", middleware.RequirePermission("create_roles"), container.RoleHandler.Create)
	roleGroup.Post("/:id/permissions", middleware.RequirePermission("edit_roles"), container.RoleHandler.AssignPermission)
	roleGroup.Put("/:id", middleware.RequirePermission("edit_roles"), container.RoleHandler.Update)
	roleGroup.Delete("/:id", middleware.RequirePermission("delete_roles"), container.RoleHandler.Delete)

	// =====================
	// Permission Routes
	// =====================
	permissionGroup := protected.Group("/permission")
	permissionGroup.Get("/", middleware.RequirePermission("view_permissions"), container.PermissionHandler.GetAll)
	permissionGroup.Post("/add", middleware.RequirePermission("create_permissions"), container.PermissionHandler.Create)
	permissionGroup.Put("/:id", middleware.RequirePermission("edit_permissions"), container.PermissionHandler.Update)
	permissionGroup.Delete("/:id", middleware.RequirePermission("delete_permissions"), container.PermissionHandler.Delete)
}
