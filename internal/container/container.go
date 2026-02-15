package container

import (
	"database/sql"
	"gpt/config"
	"gpt/internal/delivery/http/auth"
	"gpt/internal/delivery/http/handler"
	"gpt/internal/domain"
	"gpt/internal/infrastructure/repository"
	"gpt/internal/usecase"
)

type Container struct {
	AuthHandler  *handler.AuthHandler
	RoleHandler  *handler.RolesHandler
	TokenService domain.TokenService
}

func NewContainer(db *sql.DB, cfg *config.Config) *Container {
	jwtService := auth.NewJWTService(cfg.JWTSecret, cfg.JWTAccessExpiry, cfg.JWTRefreshExpiry)

	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)

	authUsecase := usecase.NewAuthUsecase(userRepo, jwtService)
	roleUsecase := usecase.NewRoleUsecase(roleRepo)

	return &Container{
		AuthHandler:  handler.NewAuthHandler(authUsecase),
		RoleHandler:  handler.NewRolesHandler(roleUsecase),
		TokenService: jwtService,
	}
}
