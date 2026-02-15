package dto

import "gpt/internal/domain"

type RolesRequest struct {
	Name      string `json:"name" validate:"required"`
	GuardName string `json:"guard_name" validate:"required"`
}

type RolesResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	GuardName string `json:"guard_name"`
}

func ToRolesResponse(roles *domain.Roles) *RolesResponse {
	return &RolesResponse{
		ID:        roles.ID,
		Name:      roles.Name,
		GuardName: roles.GuardName,
	}
}
