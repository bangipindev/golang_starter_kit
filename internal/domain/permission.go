package domain

import (
	"context"

	"github.com/google/uuid"
)

type Permission struct {
	ID        int64
	Name      string
	GuardName string
}

type PermissionRepository interface {
	GetAll(ctx context.Context) ([]Permission, error)
	GetByID(ctx context.Context, id string) (*Permission, error)
	GetUserByPublicID(PublicId uuid.UUID) (*User, error)
	GetByName(ctx context.Context, name string) (*Permission, error)
	Create(ctx context.Context, permission *Permission) error
	Update(ctx context.Context, permission *Permission) error
	Delete(ctx context.Context, id string) error
	GetUserPermissions(ctx context.Context, userID int64) ([]string, error)
	GetUserRoles(ctx context.Context, userID int64) ([]string, error)
}

type PermissionUseCase interface {
	GetAll(ctx context.Context) ([]Permission, error)
	GetByID(ctx context.Context, id string) (*Permission, error)
	GetUserByPublicID(PublicId uuid.UUID) (*User, error)
	Create(ctx context.Context, permission *Permission) error
	Update(ctx context.Context, permission *Permission) error
	Delete(ctx context.Context, id string) error
	HasPermission(ctx context.Context, userID int64, required string) (bool, error)
	HasRole(ctx context.Context, userID int64, requiredRole string) (bool, error)
}
