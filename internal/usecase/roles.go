package usecase

import (
	"context"
	"errors"
	"gpt/internal/domain"
)

type RolesUsecase interface {
	GetAll(ctx context.Context) ([]*domain.Roles, error)
	Create(ctx context.Context, roles *domain.Roles) error
	Update(ctx context.Context, role *domain.Roles) error
	Delete(ctx context.Context, id int64) error
}

type rolesUsecase struct {
	rolesRepo domain.RoleRepository
}

func NewRoleUsecase(repo domain.RoleRepository) RolesUsecase {
	return &rolesUsecase{
		rolesRepo: repo,
	}
}

func (s *rolesUsecase) GetAll(ctx context.Context) ([]*domain.Roles, error) {
	roles, err := s.rolesRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (s *rolesUsecase) Create(ctx context.Context, roles *domain.Roles) error {
	existing, _ := s.rolesRepo.FindByName(ctx, roles.Name)
	if existing != nil {
		return errors.New("Role Name already Exists")
	}

	return s.rolesRepo.Create(ctx, roles)
}

func (s *rolesUsecase) Update(ctx context.Context, role *domain.Roles) error {
	// cek apakah role ada berdasarkan ID
	existing, err := s.rolesRepo.FindByID(ctx, role.ID)
	if err != nil {
		return errors.New("role not found")
	}

	// optional: cek jika nama diubah dan sudah dipakai role lain
	if existing.Name != role.Name {
		duplicate, _ := s.rolesRepo.FindByName(ctx, role.Name)
		if duplicate != nil {
			return errors.New("role name already exists")
		}
	}

	return s.rolesRepo.Update(ctx, role)
}

func (s *rolesUsecase) Delete(ctx context.Context, id int64) error {
	// pastikan data ada dulu
	existing, err := s.rolesRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if existing == nil {
		return errors.New("role not found")
	}

	return s.rolesRepo.Delete(ctx, id)
}
