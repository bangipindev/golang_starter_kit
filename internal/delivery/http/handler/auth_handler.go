package handler

import (
	"gpt/internal/delivery/http/dto"
	"gpt/internal/domain"
	"gpt/internal/pkg/response"
	"gpt/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

func NewAuthHandler(authUsecase usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		authUsecase: authUsecase,
	}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, 400, "Invalid request body", err.Error())
	}

	user := &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	if err := h.authUsecase.Register(c.Context(), user); err != nil {
		return response.Error(c, 500, "Failed to register user", err.Error())
	}

	return response.Success(c, 201, "User registered successfully", user)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, 400, "Invalid request body", err.Error())
	}

	res, err := h.authUsecase.Login(c.Context(), req.Email, req.Password)
	if err != nil {
		return response.Error(c, 401, "Invalid credentials", err.Error())
	}

	user := dto.UserResponse{
		ID:    res.User.ID,
		Name:  res.User.Name,
		Email: res.User.Email,
	}

	loginResponse := dto.LoginResponse{
		User:         user,
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}

	return response.Success(c, 200, "Login successful", loginResponse)
}

func (h *AuthHandler) Profile(c *fiber.Ctx) error {
	userIDRaw := c.Locals("user_id")

	if userIDRaw == nil {
		return response.Error(c, 401, "Unauthorized", "invalid token")
	}

	userID := int64(userIDRaw.(float64)) // JWT numeric jadi float64

	user, err := h.authUsecase.GetProfile(c.Context(), userID)
	if err != nil {
		return response.Error(c, 404, "User not found", err.Error())
	}

	return response.Success(c, 200, "Profile Didapatkan", fiber.Map{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
	})
}

func (h *AuthHandler) Refresh(c *fiber.Ctx) error {
	var req struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, 400, "Invalid request", err.Error())
	}

	newAccess, err := h.authUsecase.RefreshToken(c.Context(), req.RefreshToken)
	if err != nil {
		return response.Error(c, 401, "Invalid refresh token", err.Error())
	}

	return response.Success(c, 200, "Token refreshed", fiber.Map{
		"access_token": newAccess,
	})
}
