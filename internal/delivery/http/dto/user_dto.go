package dto

import (
	"gpt/internal/domain"
	"gpt/internal/helper"

	"github.com/google/uuid"
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
	// Role     string `json:"role" validate:"required,oneof=admin user"`
}

type AuthUserResponse struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	PublicId uuid.UUID `json:"public_id"`
}

func ToAuthUserResponse(user *domain.User) *AuthUserResponse {
	return &AuthUserResponse{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		PublicId: user.PublicId,
	}
}

type UpdateUserRequest struct {
	Name     string  `json:"name" validate:"required"`
	Email    string  `json:"email" validate:"required,email"`
	Password *string `json:"password" validate:"omitempty,min=6"`
	// Role     string  `json:"role" validate:"required,oneof=admin user"`
}

type UpdateUserResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	PublicId  uuid.UUID `json:"public_id"`
	Status    int       `json:"status"`
	UpdatedAt string    `json:"updated_at"`
	// Role  domain.Role `json:"role"`
}

func ToUpdateUserResponse(user *domain.User) *UpdateUserResponse {
	return &UpdateUserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		PublicId:  user.PublicId,
		Status:    int(user.Status),
		UpdatedAt: helper.ToWIBString(user.UpdatedAt),
		// Role:  user.Role,
	}
}

type UserResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	PublicId  uuid.UUID `json:"public_id"`
	Status    int       `json:"status"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	// Role  domain.Role `json:"role"`
}

func ToUserResponse(user *domain.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		PublicId:  user.PublicId,
		Status:    int(user.Status),
		CreatedAt: helper.ToWIBString(user.CreatedAt),
		UpdatedAt: helper.ToWIBString(user.UpdatedAt),

		// Role:  user.Role,
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
	// role := domain.Role(r.Role)

	// Optional: validasi manual jika perlu
	// switch role {
	// case domain.RoleAdmin, domain.RoleUser:
	// default:
	// 	return nil, fmt.Errorf("invalid role")
	// }

	return &domain.User{
		Name:     r.Name,
		Email:    r.Email,
		Password: r.Password,
		PublicId: uuid.New(),
		Status:   domain.Aktif,
		// Role:     role,
	}, nil
}

type UserWithRolesPermissions struct {
	UserResponse
	Roles       []string `json:"roles"`
	Permissions []string `json:"permissions"`
}

type UserRolesPermissionsResponse struct {
	User UserWithRolesPermissions `json:"user"`
}
