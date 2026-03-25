package middleware

import (
	"fmt"
	"gpt/internal/domain"
	"gpt/internal/pkg/response"

	"github.com/gofiber/fiber/v2"
)

func RequirePermission(requiredPermission string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals("user").(*domain.AccessClaims)
		if !ok {
			return response.HandleError(c, response.ErrForbiddenAccess)
		}

		hasAccess := false

		// Superadmin bypass
		for _, r := range claims.Roles {
			if r == "superadmin" {
				hasAccess = true
				break
			}
		}

		if !hasAccess {
			for _, p := range claims.Permissions {
				if p == requiredPermission {
					fmt.Println("Permission found:", p)
					hasAccess = true
					break
				}
			}
		}

		if !hasAccess {
			return response.HandleError(c, response.ErrForbiddenAccess)
		}

		return c.Next()
	}
}

func RequireRole(requiredRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals("user").(*domain.AccessClaims)
		if !ok {
			return response.HandleError(c, response.ErrForbiddenAccess)
		}

		for _, r := range claims.Roles {
			if r == requiredRole || r == "superadmin" {
				return c.Next()
			}
		}

		return response.HandleError(c, response.ErrForbiddenAccess)
	}
}
