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
	AuthHandler       *handler.AuthHandler
	TokenService      domain.TokenService
	RoleHandler       *handler.RolesHandler
	UserHandler       *handler.UserHandler
	PermissionHandler *handler.PermissionHandler
	UserRepo          domain.UserRepository
}

func NewContainer(db *sql.DB, cfg *config.Config) *Container {
	jwtService := auth.NewJWTService(cfg.JWTSecret, cfg.JWTAccessExpiry, cfg.JWTRefreshExpiry)

	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	permissionRepo := repository.NewPermissionRepository(db)

	authUsecase := usecase.NewAuthUsecase(userRepo, jwtService)
	roleUsecase := usecase.NewRoleUsecase(roleRepo)
	userUsecase := usecase.NewUserUsecase(userRepo)
	permissionUsecase := usecase.NewPermissionUseCase(permissionRepo)

	return &Container{
		TokenService:      jwtService,
		UserRepo:          userRepo,
		AuthHandler:       handler.NewAuthHandler(authUsecase),
		RoleHandler:       handler.NewRolesHandler(roleUsecase),
		UserHandler:       handler.NewUserHandler(userUsecase),
		PermissionHandler: handler.NewPermissionHandler(permissionUsecase),
	}
}
