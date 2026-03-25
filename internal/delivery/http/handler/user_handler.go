package handler

import (
	"gpt/internal/delivery/http/dto"
	"gpt/internal/domain"
	"gpt/internal/pkg/response"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUsecase domain.UserUsecase
}

func NewUserHandler(userUsecase domain.UserUsecase) *UserHandler {
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

func (h *UserHandler) Update(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return response.HandleError(c, response.ErrorBadRequest)
	}

	var req dto.UpdateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return response.HandleError(c, response.ErrorBadRequest)
	}

	if err := validate.Struct(req); err != nil {
		return response.ValidationError(c, err)
	}

	user := &domain.User{
		ID:    id,
		Name:  req.Name,
		Email: req.Email,
	}

	if req.Password != nil {
		user.Password = *req.Password
	}

	updatedUser, err := h.userUsecase.Update(c.Context(), user)
	if err != nil {
		return response.HandleError(c, err)
	}

	return response.SuccessWithStatus(c, fiber.StatusOK, "User updated successfully", dto.ToUpdateUserResponse(updatedUser))
}

func (h *UserHandler) Delete(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return response.HandleError(c, response.ErrorBadRequest)
	}

	if err := h.userUsecase.Delete(c.Context(), id); err != nil {
		return response.HandleError(c, err)
	}

	return response.SuccessWithStatus(c, fiber.StatusOK, "User deleted successfully", nil)
}

func (h *UserHandler) AssignRole(c *fiber.Ctx) error {
	idParam := c.Params("id")
	userID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return response.HandleError(c, response.ErrorBadRequest)
	}

	var req struct {
		RoleID int64 `json:"role_id"`
	}
	if err := c.BodyParser(&req); err != nil {
		return response.HandleError(c, response.ErrorBadRequest)
	}

	if err := h.userUsecase.AssignRoleToUser(c.Context(), userID, req.RoleID); err != nil {
		return response.HandleError(c, err)
	}

	return response.SuccessWithStatus(c, fiber.StatusOK, "Role assigned successfully", nil)
}

func (h *UserHandler) AssignPermission(c *fiber.Ctx) error {
	idParam := c.Params("id")
	userID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return response.HandleError(c, response.ErrorBadRequest)
	}

	var req struct {
		PermissionID int64 `json:"permission_id"`
	}
	if err := c.BodyParser(&req); err != nil {
		return response.HandleError(c, response.ErrorBadRequest)
	}

	if err := h.userUsecase.AssignPermissionToUser(c.Context(), userID, req.PermissionID); err != nil {
		return response.HandleError(c, err)
	}

	return response.SuccessWithStatus(c, fiber.StatusOK, "Permission assigned successfully", nil)
}

func (h *UserHandler) GetRolesAndPermissions(c *fiber.Ctx) error {
	idParam := c.Params("id")
	userID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		return response.HandleError(c, response.ErrorBadRequest)
	}

	user, err := h.userUsecase.FindByID(c.Context(), userID)
	if err != nil {
		return response.HandleError(c, response.ErrorBadRequest)
	}

	roles, permissions, err := h.userUsecase.GetRolesAndPermissions(c.Context(), userID)
	if err != nil {
		return response.HandleError(c, err)
	}

	res := dto.UserRolesPermissionsResponse{
		User: dto.UserWithRolesPermissions{
			UserResponse: dto.ToUserResponse(user),
			Roles:        roles,
			Permissions:  permissions,
		},
	}

	return response.SuccessWithStatus(c, fiber.StatusOK, "Successfully fetched roles and permissions", res)
}
