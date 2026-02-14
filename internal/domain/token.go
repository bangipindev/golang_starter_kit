package domain

type TokenService interface {
	GenerateAccessToken(user *User) (string, error)
	GenerateRefreshToken(user *User) (string, error)
	ParseAccessToken(token string) (*AccessClaims, error)
	ParseRefreshToken(token string) (int64, error)
}
