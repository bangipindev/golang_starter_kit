package dto

import "gpt/internal/domain"

type RolesRequest struct {
	Name      string `json:"name" validate:"required"`
	GuardName string `json:"guard_name" validate:"required"`
}

type RolesResponse struct {
	Name      string `json:"name"`
	GuardName string `json:"guard_name"`
}

func ToRolesResponse(roles *domain.Roles) *RolesResponse {
	return &RolesResponse{
		Name:      roles.Name,
		GuardName: roles.GuardName,
	}
}
