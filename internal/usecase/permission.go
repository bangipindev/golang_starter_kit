package usecase

import (
	"context"
	"gpt/internal/domain"
	"gpt/internal/pkg/response"
	"strings"

	"github.com/google/uuid"
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

func (u *PermissionUseCase) GetUserByPublicID(PublicId uuid.UUID) (*domain.User, error) {
	return u.permissionRepo.GetUserByPublicID(PublicId)
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

func (u *PermissionUseCase) HasPermission(ctx context.Context, userID int64, required string) (bool, error) {
	// ambil role
	roles, err := u.permissionRepo.GetUserRoles(ctx, userID)
	if err != nil {
		return false, err
	}

	// superadmin bypass
	for _, r := range roles {
		if r == "superadmin" {
			return true, nil
		}
	}

	// ambil permission
	perms, err := u.permissionRepo.GetUserPermissions(ctx, userID)
	if err != nil {
		return false, err
	}

	for _, p := range perms {
		// exact match
		if p == required {
			return true, nil
		}

		// wildcard support: user.*
		if strings.HasSuffix(p, "*") {
			prefix := strings.TrimSuffix(p, "*")
			if strings.HasPrefix(required, prefix) {
				return true, nil
			}
		}
	}

	return false, nil
}

func (u *PermissionUseCase) HasRole(ctx context.Context, userID int64, requiredRole string) (bool, error) {
	roles, err := u.permissionRepo.GetUserRoles(ctx, userID)
	if err != nil {
		return false, err
	}

	for _, r := range roles {
		if r == requiredRole || r == "superadmin" {
			return true, nil
		}
	}

	return false, nil
}
