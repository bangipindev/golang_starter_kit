package handler

import (
	"gpt/internal/delivery/http/dto"
	"gpt/internal/domain"
	"gpt/internal/pkg/response"
	"gpt/internal/usecase"
	"strconv"

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

func (h *RolesHandler) GetAll(c *fiber.Ctx) error {
	ctx := c.Context()

	roles, err := h.rolesUsecase.GetAll(ctx)
	if err != nil {
		return response.Error(c, 500, "Internal Server Error", err.Error())
	}

	return response.Success(c, 200, "Successfully", roles)
}

func (h *RolesHandler) Create(c *fiber.Ctx) error {
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

	if err := h.rolesUsecase.Create(c.Context(), roles); err != nil {
		return response.Error(c, 500, "Failed to register roles", err.Error())
	}

	return response.Success(c, 201, "Role registered successfully", roles)
}

func (h *RolesHandler) Update(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return response.Error(c, 400, "Invalid role ID", err.Error())
	}

	var req dto.RolesRequest

	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, 400, "Invalid request body", err.Error())
	}

	if err := validate.Struct(req); err != nil {
		return response.Error(c, 400, "Validation failed", err.Error())
	}

	role := &domain.Roles{
		ID:        id,
		Name:      req.Name,
		GuardName: req.GuardName,
	}

	if err := h.rolesUsecase.Update(c.Context(), role); err != nil {
		return response.Error(c, 400, "Failed to update role", err.Error())
	}

	return response.Success(c, 200, "Role updated successfully", role)
}

func (h *RolesHandler) Delete(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return response.Error(c, 400, "Invalid role ID", err.Error())
	}

	if err := h.rolesUsecase.Delete(c.Context(), id); err != nil {
		return response.Error(c, 404, "Role not found", err.Error())
	}

	return response.Success(c, 200, "Role deleted successfully", nil)
}
