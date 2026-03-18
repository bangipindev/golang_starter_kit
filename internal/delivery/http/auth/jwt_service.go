package auth

import (
	"gpt/internal/domain"
	"gpt/internal/utils/token"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTService struct {
	secret        string
	accessExpiry  time.Duration
	refreshExpiry time.Duration
}

func NewJWTService(secret string, accessExpiry, refreshExpiry time.Duration) *JWTService {
	return &JWTService{
		secret:        secret,
		accessExpiry:  accessExpiry,
		refreshExpiry: refreshExpiry,
	}
}

func (j *JWTService) GenerateAccessToken(user *domain.User) (string, error) {
	claims := token.JWTAccessClaims{
		PublicId: user.PublicId,
		Name:     user.Name,
		// Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.accessExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tkn.SignedString([]byte(j.secret))
}

func (j *JWTService) GenerateRefreshToken(user *domain.User) (string, error) {
	claims := token.JWTRefreshClaims{
		PublicId: user.PublicId,
		Type:     "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.refreshExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secret))
}

func (j *JWTService) ParseRefreshToken(tokenString string) (uuid.UUID, error) {
	parsedToken, err := jwt.ParseWithClaims(
		tokenString,
		&token.JWTRefreshClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(j.secret), nil
		},
	)
	if err != nil || !parsedToken.Valid {
		return uuid.Nil, err
	}

	claims, ok := parsedToken.Claims.(*token.JWTRefreshClaims)
	if !ok {
		return uuid.Nil, jwt.ErrTokenInvalidClaims
	}

	if claims.Type != "refresh" {
		return uuid.Nil, jwt.ErrTokenInvalidClaims
	}

	return claims.PublicId, nil
}

func (j *JWTService) ParseAccessToken(tokenString string) (*domain.AccessClaims, error) {
	jwtClaims := &token.JWTAccessClaims{}

	tkn, err := jwt.ParseWithClaims(
		tokenString,
		jwtClaims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(j.secret), nil
		},
	)

	if err != nil || !tkn.Valid {
		return nil, err
	}

	// convert ke domain claims
	return &domain.AccessClaims{
		PublicId: jwtClaims.PublicId,
		Name:     jwtClaims.Name,
	}, nil
}
