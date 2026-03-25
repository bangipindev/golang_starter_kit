package container

import (
	"database/sql"
	"gpt/config"
	"gpt/internal/delivery/http/handler"
	"gpt/internal/domain"
	"gpt/internal/infrastructure/repository"
	"gpt/internal/usecase"
	"gpt/internal/utils/auth"
)

type Container struct {
	TokenService      domain.TokenService
	UserRepo          domain.UserRepository
	PermitRbac        domain.PermissionUseCase
	AuthHandler       *handler.AuthHandler
	RoleHandler       *handler.RolesHandler
	UserHandler       *handler.UserHandler
	PermissionHandler *handler.PermissionHandler
}

func NewContainer(db *sql.DB, cfg *config.Config) *Container {
	jwtService := auth.NewJWTService(cfg.JWTSecret, cfg.JWTAccessExpiry, cfg.JWTRefreshExpiry)

	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)
	permissionRepo := repository.NewPermissionRepository(db)

	authUsecase := usecase.NewAuthUsecase(userRepo, jwtService)
	roleUsecase := usecase.NewRoleUsecase(roleRepo)
	userUsecase := usecase.NewUserUsecase(userRepo)
	permitRbac := usecase.NewPermissionUseCase(permissionRepo)
	permissionUsecase := usecase.NewPermissionUseCase(permissionRepo)

	return &Container{
		TokenService:      jwtService,
		UserRepo:          userRepo,
		PermitRbac:        permitRbac,
		AuthHandler:       handler.NewAuthHandler(authUsecase),
		RoleHandler:       handler.NewRolesHandler(roleUsecase),
		UserHandler:       handler.NewUserHandler(userUsecase),
		PermissionHandler: handler.NewPermissionHandler(permissionUsecase),
	}
}
