package handler

import (
	"gpt/internal/delivery/http/dto"
	"gpt/internal/domain"
	"gpt/internal/pkg/response"
	"gpt/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

type RolesHandler struct {
	rolesUsecase usecase.RolesUsecase
}

func NewRolesHandler(rolesUsecase usecase.RolesUsecase) *RolesHandler {
	return &RolesHandler{
		rolesUsecase: rolesUsecase,
	}
}

func (h *RolesHandler) GetRoles(c *fiber.Ctx) error {
	ctx := c.Context()

	roles, err := h.rolesUsecase.GetRoles(ctx)
	if err != nil {
		return response.Error(c, 500, "Internal Server Error", err.Error())
	}

	return response.Success(c, 200, "Successfully", roles)
}

func (h *RolesHandler) Add(c *fiber.Ctx) error {
	var req dto.RolesRequest

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, 400, "Invalid request body", err.Error())
	}

	// Validasi struct
	if err := validate.Struct(req); err != nil {
		// Bisa return error validasi dengan pesan detail
		return response.Error(c, 400, "Validation failed", err.Error())
	}

	roles := &domain.Roles{
		Name:      req.Name,
		GuardName: req.GuardName,
	}

	if err := h.rolesUsecase.Add(c.Context(), roles); err != nil {
		return response.Error(c, 500, "Failed to register roles", err.Error())
	}

	return response.Success(c, 201, "Role registered successfully", roles)
}
