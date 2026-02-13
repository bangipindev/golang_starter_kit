package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func JWTProtected(secret string) fiber.Handler {
    return func(c *fiber.Ctx) error {

        tokenString := strings.TrimPrefix(
            c.Get("Authorization"),
            "Bearer ",
        )

        claims := jwt.MapClaims{}

        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return []byte(secret), nil
        })

        if err != nil || !token.Valid {
            return c.Status(401).JSON(fiber.Map{"error": "invalid token"})
        }

        c.Locals("user_id", claims["user_id"])
		c.Locals("token", tokenString)
        return c.Next()
    }
}
