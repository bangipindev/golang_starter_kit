package domain

import "context"

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
}

type RoleRepository interface {
	Create(ctx context.Context, roles *Roles) error
	FindByName(ctx context.Context, name string) (*Roles, error)
	GetRoles(ctx context.Context) ([]*Roles, error)
}
