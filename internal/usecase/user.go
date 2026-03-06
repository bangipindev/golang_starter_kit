package usecase

import (
	"context"
	"gpt/internal/domain"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	GetAll(ctx context.Context) ([]*domain.User, error)
	Create(ctx context.Context, user *domain.User) error
}

type userUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: repo,
	}
}

func (s *userUsecase) GetAll(ctx context.Context) ([]*domain.User, error) {
	user, err := s.userRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *userUsecase) Create(ctx context.Context, user *domain.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	return u.userRepo.Create(ctx, user)
}
