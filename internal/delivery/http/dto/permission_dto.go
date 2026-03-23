package dto

import (
	"gpt/internal/domain"
	"strconv"
)

type PermissionRequest struct {
	Name      string `json:"name" validate:"required"`
	GuardName string `json:"guard_name" validate:"required"`
}

type PermissionResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	GuardName string `json:"guard_name"`
}

func ToPermissionResponse(permission domain.Permission) PermissionResponse {
	return PermissionResponse{
		ID:        strconv.FormatInt(permission.ID, 10),
		Name:      permission.Name,
		GuardName: permission.GuardName,
	}
}

func ToPermissionResponseList(permissions []domain.Permission) []PermissionResponse {
	var responses []PermissionResponse
	for _, permission := range permissions {
		responses = append(responses, ToPermissionResponse(permission))
	}
	if responses == nil {
		responses = make([]PermissionResponse, 0)
	}
	return responses
}
