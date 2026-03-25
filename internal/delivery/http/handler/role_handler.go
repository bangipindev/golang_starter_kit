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
		return response.HandleError(c, err)
	}

	roleResponse := dto.ToRolesResponseList(roles)
	return response.SuccessWithStatus(c, fiber.StatusOK, "Successfully", roleResponse)
}

func (h *RolesHandler) Create(c *fiber.Ctx) error {
	var req dto.RolesRequest

	if err := c.BodyParser(&req); err != nil {
		return response.HandleError(c, response.ErrorBadRequest)
	}

	if err := validate.Struct(req); err != nil {
		return response.ValidationError(c, err)
	}

	roles := &domain.Roles{
		Name:      req.Name,
		GuardName: req.GuardName,
	}

	if err := h.rolesUsecase.Create(c.Context(), roles); err != nil {
		return response.HandleError(c, err)
	}

	roleRes := dto.ToRolesResponse(roles)
	return response.SuccessWithStatus(c, fiber.StatusCreated, "Role registered successfully", roleRes)
}

func (h *RolesHandler) Update(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return response.HandleError(c, response.ErrorBadRequest)
	}

	var req dto.RolesRequest

	if err := c.BodyParser(&req); err != nil {
		return response.HandleError(c, response.ErrorBadRequest)
	}

	if err := validate.Struct(req); err != nil {
		return response.ValidationError(c, err)
	}

	role := &domain.Roles{
		ID:        id,
		Name:      req.Name,
		GuardName: req.GuardName,
	}

	if err := h.rolesUsecase.Update(c.Context(), role); err != nil {
		return response.HandleError(c, err)
	}

	return response.SuccessWithStatus(c, fiber.StatusOK, "Role updated successfully", role)
}

func (h *RolesHandler) Delete(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return response.HandleError(c, response.ErrorBadRequest)
	}

	if err := h.rolesUsecase.Delete(c.Context(), id); err != nil {
		return response.HandleError(c, err)
	}

	return response.SuccessWithStatus(c, fiber.StatusOK, "Role deleted successfully", nil)
}

func (h *RolesHandler) AssignPermission(c *fiber.Ctx) error {
	idParam := c.Params("id")
	roleID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return response.HandleError(c, response.ErrorBadRequest)
	}

	var req struct {
		PermissionID int64 `json:"permission_id"`
	}
	if err := c.BodyParser(&req); err != nil {
		return response.HandleError(c, response.ErrorBadRequest)
	}

	if err := h.rolesUsecase.AssignPermissionToRole(c.Context(), roleID, req.PermissionID); err != nil {
		return response.HandleError(c, err)
	}

	return response.SuccessWithStatus(c, fiber.StatusOK, "Permission assigned successfully", nil)
}
