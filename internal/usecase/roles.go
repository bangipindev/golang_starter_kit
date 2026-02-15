package usecase

import (
	"context"
	"errors"
	"gpt/internal/domain"
)

type RolesUsecase interface {
	Add(ctx context.Context, roles *domain.Roles) error
	GetRoles(ctx context.Context) ([]*domain.Roles, error)
}

type rolesUsecase struct {
	rolesRepo domain.RoleRepository
}

func NewRoleUsecase(repo domain.RoleRepository) RolesUsecase {
	return &rolesUsecase{
		rolesRepo: repo,
	}
}

func (s *rolesUsecase) Add(ctx context.Context, roles *domain.Roles) error {
	existing, _ := s.rolesRepo.FindByName(ctx, roles.Name)
	if existing != nil {
		return errors.New("Role Name already Exists")
	}

	return s.rolesRepo.Create(ctx, roles)
}

func (s *rolesUsecase) GetRoles(ctx context.Context) ([]*domain.Roles, error) {
	roles, err := s.rolesRepo.GetRoles(ctx)
	if err != nil {
		return nil, err
	}

	return roles, nil
}
