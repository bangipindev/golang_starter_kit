package domain

type AccessClaims struct {
	UserID int64  `json:"user_id"`
	Name   string `json:"name"`
	Role   Role   `json:"role"`
}

type RefreshClaims struct {
	UserID int64  `json:"user_id"`
	Type   string `json:"type"`
}
