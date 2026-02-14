package token

import (
	"gpt/internal/domain"

	"github.com/golang-jwt/jwt/v5"
)

type JWTAccessClaims struct {
	UserID int64       `json:"user_id"`
	Name   string      `json:"name"`
	Role   domain.Role `json:"role"`
	jwt.RegisteredClaims
}

type JWTRefreshClaims struct {
	UserID int64  `json:"user_id"`
	Type   string `json:"type"`
	jwt.RegisteredClaims
}
