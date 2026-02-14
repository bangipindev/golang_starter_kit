package usecase

import (
	"context"
	"errors"
	"gpt/internal/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Register(ctx context.Context, user *domain.User) error
	Login(ctx context.Context, email, password string) (*LoginResponse, error)
	GetProfile(ctx context.Context, id int64) (*domain.User, error)
	RefreshToken(ctx context.Context, refreshToken string) (string, error)
}

type authUsecase struct {
	userRepo      domain.UserRepository
	secret        string
	accessExpiry  time.Duration
	refreshExpiry time.Duration
}

type LoginResponse struct {
	User         *domain.User
	AccessToken  string
	RefreshToken string
}

type AccessClaims struct {
	UserID int64  `json:"user_id"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	UserID int64  `json:"user_id"`
	Type   string `json:"type"`
	jwt.RegisteredClaims
}

func NewAuthUsecase(repo domain.UserRepository, secret string, accessExpiry time.Duration,
	refreshExpiry time.Duration,
) AuthUsecase {
	return &authUsecase{
		userRepo:      repo,
		secret:        secret,
		accessExpiry:  accessExpiry,
		refreshExpiry: refreshExpiry,
	}
}

func (s *authUsecase) Register(ctx context.Context, user *domain.User) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashed)
	return s.userRepo.Create(ctx, user)
}

func (s *authUsecase) Login(ctx context.Context, email, password string) (*LoginResponse, error) {
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	// Access Token (1 day)
	claims := AccessClaims{
		UserID: user.ID,
		Name:   user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.accessExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, _ := accessToken.SignedString([]byte(s.secret))

	// Refresh Token (7 days)
	refreshClaims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(s.refreshExpiry).Unix(),
		"type":    "refresh",
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, _ := refreshToken.SignedString([]byte(s.secret))

	user.Password = ""

	return &LoginResponse{
		User:         user,
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

func (s *authUsecase) GetProfile(ctx context.Context, userID int64) (*domain.User, error) {
	return s.userRepo.FindByID(ctx, userID)
}

func (s *authUsecase) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	token, err := jwt.ParseWithClaims(
		refreshToken,
		&RefreshClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(s.secret), nil
		},
	)
	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(*RefreshClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid refresh token")
	}

	// Pastikan ini refresh token
	if claims.Type != "refresh" {
		return "", errors.New("invalid token type")
	}

	userID := claims.UserID

	// Optional: cek user masih ada
	_, err = s.userRepo.FindByID(ctx, userID)
	if err != nil {
		return "", errors.New("user not found")
	}

	// Generate new access token
	return s.generateAccessToken(userID)
}

func (s *authUsecase) generateAccessToken(userID int64) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.accessExpiry)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, struct {
		UserID int64 `json:"user_id"`
		jwt.RegisteredClaims
	}{
		UserID:           userID,
		RegisteredClaims: claims,
	})

	return token.SignedString([]byte(s.secret))
}
