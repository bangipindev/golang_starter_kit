package handler

import (
	"gpt/internal/delivery/http/dto"
	"gpt/internal/domain"
	"gpt/internal/pkg/response"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type PermissionHandler struct {
	PermissionUseCase domain.PermissionUseCase
}

func NewPermissionHandler(permissionUseCase domain.PermissionUseCase) *PermissionHandler {
	return &PermissionHandler{PermissionUseCase: permissionUseCase}
}

func (h *PermissionHandler) GetAll(c *fiber.Ctx) error {
	permissions, err := h.PermissionUseCase.GetAll(c.Context())
	if err != nil {
		return err
	}
	permissionResponse := dto.ToPermissionResponseList(permissions)
	return response.SuccessWithStatus(c, fiber.StatusOK, "Successfully", permissionResponse)
}

func (h *PermissionHandler) Create(c *fiber.Ctx) error {
	var permission dto.PermissionRequest
	if err := c.BodyParser(&permission); err != nil {
		return response.HandleError(c, err)
	}
	if err := validate.Struct(permission); err != nil {
		return response.ValidationError(c, err)
	}
	permissionDomain := domain.Permission{
		Name:      strings.ToLower(permission.Name),
		GuardName: strings.ToLower(permission.GuardName),
	}
	if err := h.PermissionUseCase.Create(c.Context(), &permissionDomain); err != nil {
		return response.HandleError(c, err)
	}
	permissionResponse := dto.ToPermissionResponse(permissionDomain)
	return response.SuccessWithStatus(c, fiber.StatusCreated, "Permission registered successfully", permissionResponse)
}

func (h *PermissionHandler) Update(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return response.HandleError(c, response.ErrorBadRequest)
	}
	var permission dto.PermissionRequest
	if err := c.BodyParser(&permission); err != nil {
		return response.HandleError(c, err)
	}
	if err := validate.Struct(permission); err != nil {
		return response.ValidationError(c, err)
	}
	permissionDomain := domain.Permission{
		ID:        id,
		Name:      permission.Name,
		GuardName: permission.GuardName,
	}
	if err := h.PermissionUseCase.Update(c.Context(), &permissionDomain); err != nil {
		return response.HandleError(c, err)
	}
	permissionResponse := dto.ToPermissionResponse(permissionDomain)
	return response.SuccessWithStatus(c, fiber.StatusOK, "Permission updated successfully", permissionResponse)
}

func (h *PermissionHandler) Delete(c *fiber.Ctx) error {
	if err := h.PermissionUseCase.Delete(c.Context(), c.Params("id")); err != nil {
		return response.HandleError(c, err)
	}
	return response.SuccessWithStatus(c, fiber.StatusOK, "Permission deleted successfully", nil)
}
