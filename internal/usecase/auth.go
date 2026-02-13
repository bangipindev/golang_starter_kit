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
	Login(ctx context.Context, email, password string) (string, error)
	GetProfile(ctx context.Context, id int64) (*domain.User, error)
}

type authUsecase struct {
	userRepo domain.UserRepository
	secret   string
}

func NewAuthUsecase(repo domain.UserRepository, secret string) AuthUsecase {
	return &authUsecase{
		userRepo: repo,
		secret:   secret,
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

func (s *authUsecase) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secret))
}

func (s *authUsecase) GetProfile(ctx context.Context, userID int64) (*domain.User, error) {
	return s.userRepo.FindByID(ctx, userID)
}
