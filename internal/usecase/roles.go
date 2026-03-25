package usecase

import (
	"context"
	"gpt/internal/domain"
	"gpt/internal/pkg/response"
)

type RolesUsecase interface {
	GetAll(ctx context.Context) ([]*domain.Roles, error)
	Create(ctx context.Context, roles *domain.Roles) error
	Update(ctx context.Context, role *domain.Roles) error
	Delete(ctx context.Context, id int64) error
	AssignPermissionToRole(ctx context.Context, roleID int64, permissionID int64) error
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
		return response.ErrEmailAlreadyUsed
	}

	return s.rolesRepo.Create(ctx, roles)
}

func (s *rolesUsecase) Update(ctx context.Context, role *domain.Roles) error {

	existing, err := s.rolesRepo.FindByID(ctx, role.ID)
	if err != nil {
		return response.ErrNotFound
	}

	// optional: cek jika nama diubah dan sudah dipakai role lain
	if existing.Name != role.Name {
		duplicate, _ := s.rolesRepo.FindByName(ctx, role.Name)
		if duplicate != nil {
			return response.ErrEmailAlreadyUsed
		}
	}

	return s.rolesRepo.Update(ctx, role)
}

func (s *rolesUsecase) Delete(ctx context.Context, id int64) error {
	// pastikan data ada dulu
	existing, err := s.rolesRepo.FindByID(ctx, id)
	if err != nil {
		return response.ErrNotFound
	}

	if existing == nil {
		return response.ErrNotFound
	}

	return s.rolesRepo.Delete(ctx, id)
}

func (s *rolesUsecase) AssignPermissionToRole(ctx context.Context, roleID int64, permissionID int64) error {
	return s.rolesRepo.AssignPermissionToRole(ctx, roleID, permissionID)
}
