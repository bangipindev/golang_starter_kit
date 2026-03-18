package middleware

import (
	"gpt/internal/domain"
	"gpt/internal/pkg/response"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(tokenService domain.TokenService, userRepo domain.UserRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return response.HandleError(c, response.ErrAuthorization)
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			return response.HandleError(c, response.ErrAuthNotValid)
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := tokenService.ParseAccessToken(tokenString)
		if err != nil {
			return response.HandleError(c, response.ErrTokenExpired)
		}

		user, err := userRepo.FindByPublicID(c.Context(), claims.PublicId)
		if err != nil || user == nil {
			return response.HandleError(c, response.ErrNotFound)
		}

		// 🚨 cek status
		if user.Status != domain.Aktif {
			return response.HandleError(c, response.ErrUserInactive)
		}

		c.Locals("user", claims)

		return c.Next()
	}
}

func GetClaims(c *fiber.Ctx) (*domain.AccessClaims, error) {
	claims, ok := c.Locals("user").(*domain.AccessClaims)
	if !ok {
		return nil, fiber.ErrUnauthorized
	}
	return claims, nil
}

// func RequireRole(role domain.Role) fiber.Handler {
// 	return func(c *fiber.Ctx) error {
// 		claims, ok := c.Locals("user").(*domain.AccessClaims)
// 		if !ok {
// 			return fiber.ErrUnauthorized
// 		}

// 		if claims.Role != role {
// 			return fiber.ErrForbidden
// 		}

// 		return c.Next()
// 	}
// }
