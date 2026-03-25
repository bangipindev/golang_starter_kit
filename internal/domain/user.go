package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type StatusUser int

const (
	Aktif    StatusUser = 1
	NonAktif StatusUser = 2
)

type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	PublicId  uuid.UUID
	Status    StatusUser
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByID(ctx context.Context, id int64) (*User, error)
	FindByPublicID(ctx context.Context, publicId uuid.UUID) (*User, error)
	GetAll(ctx context.Context) ([]*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int64) error
	GetRolesAndPermissions(ctx context.Context, userID int64) ([]string, []string, error)
	AssignRoleToUser(ctx context.Context, userID int64, roleID int64) error
	AssignPermissionToUser(ctx context.Context, userID int64, permissionID int64) error
}

type UserUsecase interface {
	GetAll(ctx context.Context) ([]*User, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) (*User, error)
	Delete(ctx context.Context, id int64) error
	AssignRoleToUser(ctx context.Context, userID int64, roleID int64) error
	AssignPermissionToUser(ctx context.Context, userID int64, permissionID int64) error
	GetRolesAndPermissions(ctx context.Context, userID int64) ([]string, []string, error)
	FindByID(ctx context.Context, id int64) (*User, error)
}
