package domain

import (
	"context"

	"github.com/google/uuid"
)

type User struct {
	ID       int64
	Name     string
	Email    string
	Password string
	PublicId uuid.UUID
	Status   int
}

type UserRepository interface {
	Create(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
	FindByID(ctx context.Context, id int64) (*User, error)
	GetAll(ctx context.Context) ([]*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int64) error
}
