package usecase

import (
	"context"
	"gpt/internal/domain"
	"gpt/internal/pkg/response"
)

type PermissionUseCase struct {
	permissionRepo domain.PermissionRepository
}

func NewPermissionUseCase(permissionRepo domain.PermissionRepository) domain.PermissionUseCase {
	return &PermissionUseCase{permissionRepo: permissionRepo}
}

func (u *PermissionUseCase) GetAll(ctx context.Context) ([]domain.Permission, error) {
	return u.permissionRepo.GetAll(ctx)
}

func (u *PermissionUseCase) GetByID(ctx context.Context, id string) (*domain.Permission, error) {
	return u.permissionRepo.GetByID(ctx, id)
}

func (u *PermissionUseCase) Create(ctx context.Context, permission *domain.Permission) error {
	existing, _ := u.permissionRepo.GetByName(ctx, permission.Name)
	if existing != nil {
		return response.ErrPermissionAlreadyUsed
	}
	return u.permissionRepo.Create(ctx, permission)
}

func (u *PermissionUseCase) Update(ctx context.Context, permission *domain.Permission) error {
	return u.permissionRepo.Update(ctx, permission)
}

func (u *PermissionUseCase) Delete(ctx context.Context, id string) error {
	return u.permissionRepo.Delete(ctx, id)
}
