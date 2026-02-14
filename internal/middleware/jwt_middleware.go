package middleware

import (
	"gpt/internal/domain"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(tokenService domain.TokenService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return fiber.ErrUnauthorized
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := tokenService.ParseAccessToken(tokenString)
		if err != nil {
			return fiber.ErrUnauthorized
		}

		c.Locals("user", claims)

		return c.Next()
	}
}

func RequireRole(role domain.Role) fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims, ok := c.Locals("user").(*domain.AccessClaims)
		if !ok {
			return fiber.ErrUnauthorized
		}

		if claims.Role != role {
			return fiber.ErrForbidden
		}

		return c.Next()
	}
}
