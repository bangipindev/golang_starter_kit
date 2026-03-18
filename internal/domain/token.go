package domain

import "github.com/google/uuid"

type TokenService interface {
	GenerateAccessToken(user *User) (string, error)
	GenerateRefreshToken(user *User) (string, error)
	ParseAccessToken(token string) (*AccessClaims, error)
	ParseRefreshToken(token string) (uuid.UUID, error)
}
