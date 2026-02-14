package auth

import (
	"gpt/internal/auth/token"
	"gpt/internal/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
		UserID: user.ID,
		Name:   user.Name,
		Role:   user.Role,
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
		UserID: user.ID,
		Type:   "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.refreshExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secret))
}

func (j *JWTService) ParseRefreshToken(tokenString string) (int64, error) {
	parsedToken, err := jwt.ParseWithClaims(
		tokenString,
		&token.JWTRefreshClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(j.secret), nil
		},
	)
	if err != nil || !parsedToken.Valid {
		return 0, err
	}

	claims, ok := parsedToken.Claims.(*token.JWTRefreshClaims)
	if !ok {
		return 0, jwt.ErrTokenInvalidClaims
	}

	if claims.Type != "refresh" {
		return 0, jwt.ErrTokenInvalidClaims
	}

	return claims.UserID, nil
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
		UserID: jwtClaims.UserID,
		Name:   jwtClaims.Name,
		Role:   jwtClaims.Role,
	}, nil
}
