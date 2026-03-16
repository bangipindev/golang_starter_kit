package dto

import (
	"fmt"
	"gpt/internal/domain"
)

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
type UserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Role     string `json:"role" validate:"required,oneof=admin user"`
}

type AuthUserResponse struct {
	ID    int64       `json:"id"`
	Name  string      `json:"name"`
	Email string      `json:"email"`
	Role  domain.Role `json:"role"`
}

func ToAuthUserResponse(user *domain.User) *AuthUserResponse {
	return &AuthUserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}
}

type UserResponse struct {
	ID    int64       `json:"id"`
	Name  string      `json:"name"`
	Email string      `json:"email"`
	Role  domain.Role `json:"role"`
}

func ToUserResponse(user *domain.User) UserResponse {
	return UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}
}

func ToUserResponseList(users []*domain.User) []UserResponse {
	result := make([]UserResponse, 0, len(users))

	for _, u := range users {
		result = append(result, ToUserResponse(u))
	}

	return result
}

func (r *UserRequest) ToDomain() (*domain.User, error) {
	role := domain.Role(r.Role)

	// Optional: validasi manual jika perlu
	switch role {
	case domain.RoleAdmin, domain.RoleUser:
	default:
		return nil, fmt.Errorf("invalid role")
	}

	return &domain.User{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password, // nanti di-hash di usecase
		Role:     role,
	}, nil
}
