package middleware

import (
	"gpt/internal/domain"
	"gpt/internal/pkg/response"

	"github.com/gofiber/fiber/v2"
)

// func RequirePermission(requiredPermission string) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		claims, ok := c.Locals("user").(*domain.AccessClaims)
// 		if !ok {
// 			return response.HandleError(c, response.ErrForbiddenAccess)
// 		}

// 		hasAccess := false

// 		// Superadmin bypass
// 		for _, r := range claims.Roles {
// 			if r == "superadmin" {
// 				hasAccess = true
// 				break
// 			}
// 		}

// 		if !hasAccess {
// 			for _, p := range claims.Permissions {
// 				if p == requiredPermission {
// 					hasAccess = true
// 					break
// 				}
// 			}
// 		}

// 		if !hasAccess {
// 			return response.HandleError(c, response.ErrForbiddenAccess)
// 		}

// 		return c.Next()
// 	}
// }

func RequirePermission(permission string, uc domain.PermissionUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals("user").(*domain.AccessClaims)
		if !ok {
			return response.HandleError(c, response.ErrForbiddenAccess)
		}

		user, err := uc.GetUserByPublicID(claims.PublicId)
		if err != nil || user == nil {
			return response.HandleError(c, response.ErrNotFound)
		}

		userID := user.ID

		has, err := uc.HasPermission(c.Context(), int64(userID), permission)
		if err != nil {
			return response.HandleError(c, err)
		}

		if !has {
			return response.HandleError(c, response.ErrForbiddenAccess)
		}

		return c.Next()
	}
}

func RequireRole(requiredRole string, uc domain.PermissionUseCase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals("user").(*domain.AccessClaims)

		if !ok {
			return response.HandleError(c, response.ErrForbiddenAccess)
		}

		user, err := uc.GetUserByPublicID(claims.PublicId)
		if err != nil || user == nil {
			return response.HandleError(c, response.ErrNotFound)
		}

		userID := user.ID

		has, err := uc.HasRole(c.Context(), userID, requiredRole)
		if err != nil {
			return response.HandleError(c, err)
		}

		if !has {
			return response.HandleError(c, response.ErrForbiddenAccess)
		}

		return c.Next()
	}
}
