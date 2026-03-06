package handler

import (
	"gpt/internal/delivery/http/dto"
	"gpt/internal/pkg/response"
	"gpt/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{
		userUsecase: userUsecase,
	}
}

func (h *UserHandler) GetAll(c *fiber.Ctx) error {
	ctx := c.Context()

	users, err := h.userUsecase.GetAll(ctx)
	if err != nil {
		return response.Error(c, 500, "Internal Server Error", err.Error())
	}

	userResponses := dto.ToUserResponseList(users)
	return response.Success(c, 200, "Successfully", userResponses)
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	var req dto.UserRequest

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, 400, "Invalid request", err.Error())
	}

	// Validasi struct (pakai validator kalau ada)
	if err := validate.Struct(req); err != nil {
		return response.Error(c, 400, "Validation error", err.Error())
	}

	user, err := req.ToDomain()
	if err != nil {
		return response.Error(c, 400, "Invalid role", err.Error())
	}

	if err := h.userUsecase.Create(c.Context(), user); err != nil {
		return response.Error(c, 500, "Failed to create user", err.Error())
	}

	return response.Success(c, 201, "User created", nil)
}
