package usecase

import (
	"context"
	"gpt/internal/domain"
	"gpt/internal/pkg/response"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	GetAll(ctx context.Context) ([]*domain.User, error)
	Create(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, user *domain.User) error
	Delete(ctx context.Context, id int64) error
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
	existing, err := u.userRepo.FindByEmail(ctx, user.Email)
	if err != nil || existing == nil {
		return response.ErrEmailAlreadyUsed
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	return u.userRepo.Create(ctx, user)
}

func (u *userUsecase) Update(ctx context.Context, req *domain.User) error {
	existing, err := u.userRepo.FindByID(ctx, req.ID)
	if err != nil || existing == nil {
		return response.ErrNotFound
	}

	existing.Name = req.Name
	existing.Email = req.Email
	if req.Role != "" {
		existing.Role = req.Role
	}

	if req.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		existing.Password = string(hashedPassword)
	}

	return u.userRepo.Update(ctx, existing)
}

func (u *userUsecase) Delete(ctx context.Context, id int64) error {
	existing, err := u.userRepo.FindByID(ctx, id)
	if err != nil || existing == nil {
		return response.ErrNotFound
	}
	return u.userRepo.Delete(ctx, id)
}
