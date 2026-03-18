package handler

import (
	"gpt/internal/delivery/http/dto"
	"gpt/internal/domain"
	"gpt/internal/pkg/response"
	"gpt/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: authUsecase,
	}
}

var validate = validator.New()

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return response.HandleError(c, response.ErrorBadRequest)
	}

	if err := validate.Struct(req); err != nil {
		return response.ValidationError(c, err)
	}

	user := &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		PublicId: uuid.New(),
		Status:   domain.Aktif,
	}

	if err := h.authUsecase.Register(c.Context(), user); err != nil {
		return response.HandleError(c, err)
	}

	return response.SuccessWithStatus(c, fiber.StatusCreated, "User registered successfully", dto.ToAuthUserResponse(user))
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return response.HandleError(c, response.ErrorBadRequest)
	}

	// validasi request
	if err := validate.Struct(req); err != nil {
		return response.ValidationError(c, err)
	}

	res, err := h.authUsecase.Login(c.Context(), req.Email, req.Password)
	if err != nil {
		return response.HandleError(c, err)
	}

	user := dto.AuthUserResponse{
		Name:     res.User.Name,
		Email:    res.User.Email,
		PublicId: res.User.PublicId,
	}

	loginResponse := dto.LoginResponse{
		User:         user,
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}

	return response.Success(c, "Login successful", loginResponse)
}

func (h *AuthHandler) Profile(c *fiber.Ctx) error {
	claims := c.Locals("user").(*domain.AccessClaims)

	user, err := h.authUsecase.GetProfile(c.Context(), claims.PublicId)
	if err != nil {
		return response.HandleError(c, response.ErrNotFound)
	}

	return response.Success(c, "Profile retrieved successfully", dto.ToAuthUserResponse(user))
}

func (h *AuthHandler) Refresh(c *fiber.Ctx) error {
	var req dto.RefreshTokenRequest

	if err := c.BodyParser(&req); err != nil {
		return response.HandleError(c, response.ErrorBadRequest)
	}

	// validasi request
	if err := validate.Struct(req); err != nil {
		return response.ValidationError(c, err)
	}

	newAccess, err := h.authUsecase.RefreshToken(c.Context(), req.RefreshToken)
	if err != nil {
		return response.HandleError(c, err)
	}

	return response.Success(c, "Token refreshed", fiber.Map{
		"access_token": newAccess,
	})
}
