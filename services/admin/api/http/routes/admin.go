package routes

import (
	"admin/api/http/handlers"
	"admin/api/http/services"
	"admin/app"
	"admin/config"
	"context"

	"github.com/labstack/echo"
)

func RegisterAdminRoutes(ctx context.Context, serverGroup *echo.Group, appContainer app.App, cfg config.ServerConfig) {
	s := services.NewAdminService(appContainer.AdminService(ctx))
	g := serverGroup.Group("/admins")
	g.POST("/block/users/:user_id", handlers.BlockUser(s))
}
