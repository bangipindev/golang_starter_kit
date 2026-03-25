package domain

import "github.com/google/uuid"

type AccessClaims struct {
	PublicId    uuid.UUID `json:"public_id"`
	Name        string    `json:"name"`
	Roles       []string  `json:"roles"`
	Permissions []string  `json:"permissions"`
}

type RefreshClaims struct {
	UserID int64  `json:"user_id"`
	Type   string `json:"type"`
}
