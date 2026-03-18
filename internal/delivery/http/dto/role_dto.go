package dto

import (
	"gpt/internal/domain"
	"gpt/internal/helper"
)

type RolesRequest struct {
	Name      string `json:"name" validate:"required"`
	GuardName string `json:"guard_name" validate:"required"`
}

type RolesResponse struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	GuardName string `json:"guard_name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func ToRolesResponse(roles *domain.Roles) RolesResponse {
	return RolesResponse{
		ID:        roles.ID,
		Name:      roles.Name,
		GuardName: roles.GuardName,
		CreatedAt: helper.ToWIBString(roles.CreatedAt),
		UpdatedAt: helper.ToWIBString(roles.UpdatedAt),
	}
}

func ToRolesResponseList(roles []*domain.Roles) []RolesResponse {
	result := make([]RolesResponse, 0, len(roles))

	for _, r := range roles {
		result = append(result, ToRolesResponse(r))
	}

	return result
}
