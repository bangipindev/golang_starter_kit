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
		return response.HandleError(c, err)
	}

	userResponses := dto.ToUserResponseList(users)
	return response.SuccessWithStatus(c, fiber.StatusOK, "Successfully", userResponses)
}

func (h *UserHandler) Create(c *fiber.Ctx) error {
	var req dto.UserRequest

	if err := c.BodyParser(&req); err != nil {
		return response.HandleError(c, response.ErrorBadRequest)
	}

	// Validasi struct (pakai validator kalau ada)
	if err := validate.Struct(req); err != nil {
		return response.ValidationError(c, err)
	}

	user, err := req.ToDomain()
	if err != nil {
		return response.HandleError(c, response.ErrorBadRequest)
	}

	if err := h.userUsecase.Create(c.Context(), user); err != nil {
		return response.HandleError(c, err)
	}

	return response.SuccessWithStatus(c, fiber.StatusCreated, "User created", nil)
}
