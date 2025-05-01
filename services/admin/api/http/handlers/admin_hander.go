package handlers

import (
	"admin/api/http/services"
	"admin/api/http/types"
	"net/http"
	"strconv"

	"github.com/goli-nababa/golibaba-backend/common"
	"github.com/labstack/echo"
)

type AdminHandler struct {
	adminService *services.AdminService
}

func NewAdminHandler(adminService *services.AdminService) *AdminHandler {
	return &AdminHandler{adminService: adminService}
}

func (h *AdminHandler) BlockEntity(c echo.Context) error {
	var req types.BlockEntityRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	if req.EntityID == "" || req.EntityType == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Both entity_id and entity_type are required",
		})
	}

	err := h.adminService.BlockEntity(c.Request().Context(), req.EntityID, req.EntityType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Entity blocked successfully",
	})
}

func (h *AdminHandler) UnblockEntity(c echo.Context) error {
	var req types.UnblockEntityRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	if req.EntityID == "" || req.EntityType == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Both entity_id and entity_type are required",
		})
	}

	err := h.adminService.UnblockEntity(c.Request().Context(), req.EntityID, req.EntityType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Entity unblocked successfully",
	})
}

func (h *AdminHandler) AssignRoleToUser(c echo.Context) error {
	var req types.AssignRoleRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request parameters"})
	}

	if req.Role == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "role is required",
		})
	}

	err = h.adminService.AssignRole(c.Request().Context(), uint(userId), req.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Role assigned successfully",
	})
}

func (h *AdminHandler) CancelRoleFromUser(c echo.Context) error {
	var req types.CancelRoleRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request parameters"})
	}

	if req.Role == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "role is required",
		})
	}

	err = h.adminService.CancelRole(c.Request().Context(), uint(userId), req.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Role canceled successfully",
	})
}

func (h *AdminHandler) AssignPermissionToRole(c echo.Context) error {
	var req types.AssignPermissionToRoleRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request parameters"})
	}

	if req.Role == "" || len(req.Permissions) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "role and permissions are required",
		})
	}

	err = h.adminService.AssignPermissionToRole(c.Request().Context(), uint(userId), req.Role, req.Permissions)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Permissions assigned successfully",
	})
}

// RevokePermissionFromRole revokes permissions from a role
func (h *AdminHandler) RevokePermissionFromRole(c echo.Context) error {
	var req types.RevokePermissionFromRoleRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request parameters"})
	}

	if req.Role == "" || len(req.Permissions) == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "role and permissions are required",
		})
	}

	err = h.adminService.RevokePermissionFromRole(c.Request().Context(), uint(userId), req.Role, req.Permissions)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Permissions revoked successfully",
	})
}

func (h *AdminHandler) PublishStatement(c echo.Context) error {
	var req types.PublishStatementRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	if req.Action == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "action is required",
		})
	}

	userIds := []common.UserID{}
	for _, userId := range req.UserIds {
		userIds = append(userIds, common.UserID(userId))
	}

	err := h.adminService.PublishStatement(c.Request().Context(), userIds, req.Action, req.Permissions)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Statement published successfully",
	})
}

func (h *AdminHandler) CancelStatement(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request parameters"})
	}

	statementID, err := strconv.Atoi(c.Param("statement_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request parameters"})
	}

	err = h.adminService.CancelStatement(c.Request().Context(), uint(userId), uint(statementID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Statement canceled successfully",
	})
}
