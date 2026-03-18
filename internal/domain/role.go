package domain

import (
	"context"
	"time"
)

type Role string

const (
	RoleSuperAdmin Role = "superadmin"
	RoleAdmin      Role = "admin"
	RoleUser       Role = "user"
)

type Roles struct {
	ID        int64
	Name      string
	GuardName string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RoleRepository interface {
	Create(ctx context.Context, roles *Roles) error
	FindByID(ctx context.Context, id int64) (*Roles, error)
	FindByName(ctx context.Context, name string) (*Roles, error)
	GetAll(ctx context.Context) ([]*Roles, error)
	Update(ctx context.Context, roles *Roles) error
	Delete(ctx context.Context, id int64) error
}
