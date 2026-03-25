package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Auth struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type LoginResponse struct {
	User         *User
	AccessToken  string
	RefreshToken string
}

type AuthUsecase interface {
	Register(ctx context.Context, user *User) error
	Login(ctx context.Context, email, password string) (*LoginResponse, error)
	GetProfile(ctx context.Context, public_id uuid.UUID) (*User, error)
	RefreshToken(ctx context.Context, refreshToken string) (string, error)
}
