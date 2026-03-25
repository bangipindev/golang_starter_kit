package token

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTAccessClaims struct {
	PublicId    uuid.UUID `json:"public_id"`
	Name        string    `json:"name"`
	Roles       []string  `json:"roles"`
	Permissions []string  `json:"permissions"`
	jwt.RegisteredClaims
}

type JWTRefreshClaims struct {
	PublicId uuid.UUID `json:"public_id"`
	Type     string    `json:"type"`
	jwt.RegisteredClaims
}
