package http

import (
	"admin/api/http/handlers"
	"admin/api/http/services"

	"github.com/labstack/echo"
)

func RegisterAdminRoutes(serverGroup *echo.Group, s *services.AdminService) {
	g := serverGroup.Group("/admin-panel")
	handler := handlers.NewAdminHandler(s)

	g.POST("/block", handler.BlockEntity)
	g.POST("/unblock", handler.UnblockEntity)
	g.POST("/users/:user_id/assign-role", handler.AssignRoleToUser)
	g.POST("/users/:user_id/cancel-role", handler.CancelRoleFromUser)
	g.POST("/users/:user_id/role/assign-permissions", handler.AssignPermissionToRole)
	g.POST("/users/:user_id/role/revoke-permissions", handler.RevokePermissionFromRole)
	g.POST("/users/publish-statement", handler.PublishStatement)
	g.POST("/users/:user_id/statements/:statement_id/cancel", handler.CancelStatement)

}
