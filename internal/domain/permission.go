package domain

import "context"

type Permission struct {
	ID        int64
	Name      string
	GuardName string
}

type PermissionRepository interface {
	GetAll(ctx context.Context) ([]Permission, error)
	GetByID(ctx context.Context, id string) (*Permission, error)
	GetByName(ctx context.Context, name string) (*Permission, error)
	Create(ctx context.Context, permission *Permission) error
	Update(ctx context.Context, permission *Permission) error
	Delete(ctx context.Context, id string) error
}

type PermissionUseCase interface {
	GetAll(ctx context.Context) ([]Permission, error)
	GetByID(ctx context.Context, id string) (*Permission, error)
	Create(ctx context.Context, permission *Permission) error
	Update(ctx context.Context, permission *Permission) error
	Delete(ctx context.Context, id string) error
}
