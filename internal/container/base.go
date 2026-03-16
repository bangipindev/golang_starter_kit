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
	TokenService domain.TokenService
	RoleHandler  *handler.RolesHandler
	UserHandler  *handler.UserHandler
}

func NewContainer(db *sql.DB, cfg *config.Config) *Container {
	jwtService := auth.NewJWTService(cfg.JWTSecret, cfg.JWTAccessExpiry, cfg.JWTRefreshExpiry)

	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)

	authUsecase := usecase.NewAuthUsecase(userRepo, jwtService)
	roleUsecase := usecase.NewRoleUsecase(roleRepo)
	userUsecase := usecase.NewUserUsecase(userRepo)

	return &Container{
		TokenService: jwtService,
		AuthHandler:  handler.NewAuthHandler(authUsecase),
		RoleHandler:  handler.NewRolesHandler(roleUsecase),
		UserHandler:  handler.NewUserHandler(userUsecase),
	}
}
