package usecase

import (
	"context"
	"database/sql"
	"gpt/internal/domain"
	"gpt/internal/pkg/response"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Register(ctx context.Context, user *domain.User) error
	Login(ctx context.Context, email, password string) (*LoginResponse, error)
	GetProfile(ctx context.Context, public_id uuid.UUID) (*domain.User, error)
	RefreshToken(ctx context.Context, refreshToken string) (string, error)
}

type authUsecase struct {
	userRepo domain.UserRepository
	tokenSvc domain.TokenService
}

type LoginResponse struct {
	User         *domain.User
	AccessToken  string
	RefreshToken string
}

func NewAuthUsecase(repo domain.UserRepository, tokenSvc domain.TokenService) AuthUsecase {
	return &authUsecase{
		userRepo: repo,
		tokenSvc: tokenSvc,
	}
}

func (s *authUsecase) Register(ctx context.Context, user *domain.User) error {
	existing, err := s.userRepo.FindByEmail(ctx, user.Email)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if existing != nil {
		return response.ErrEmailAlreadyUsed
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashed)
	// user.Role = domain.RoleUser

	return s.userRepo.Create(ctx, user)
}

func (s *authUsecase) Login(ctx context.Context, email, password string) (*LoginResponse, error) {
	user, err := s.userRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, response.ErrNotFound
	}

	if user.Status == domain.NonAktif {
		return nil, response.ErrUserInactive
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, response.ErrPasswordNotMatch
	}

	accessToken, err := s.tokenSvc.GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.tokenSvc.GenerateRefreshToken(user)
	if err != nil {
		return nil, err
	}

	user.Password = ""

	return &LoginResponse{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *authUsecase) GetProfile(ctx context.Context, PublicId uuid.UUID) (*domain.User, error) {
	user, err := s.userRepo.FindByPublicID(ctx, PublicId)
	if err != nil {
		return nil, response.ErrNotFound
	}

	user.Password = ""
	return user, nil
}

func (s *authUsecase) RefreshToken(ctx context.Context, refreshToken string) (string, error) {
	publicId, err := s.tokenSvc.ParseRefreshToken(refreshToken)
	if err != nil {
		return "", response.ErrorRefreshTokenInvalid
	}

	user, err := s.userRepo.FindByPublicID(ctx, publicId)
	if err != nil {
		return "", response.ErrNotFound
	}

	return s.tokenSvc.GenerateAccessToken(user)
}
