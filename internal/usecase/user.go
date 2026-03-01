package usecase

import (
	"context"
	"gpt/internal/domain"
)

type UserUsecase interface {
	GetAll(ctx context.Context) ([]*domain.User, error)
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
