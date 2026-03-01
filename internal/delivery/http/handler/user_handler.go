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
