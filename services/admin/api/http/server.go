package http

import (
	"admin/api/http/services"
	di "admin/app"
	"admin/config"
	"context"
	"fmt"

	"github.com/labstack/echo"
)

func Bootstrap(appContainer di.App, cfg config.Config) error {
	e := echo.New()
	adminService := services.NewAdminService(appContainer.AdminService(context.Background()))

	api := e.Group("/api/v1")
	RegisterAdminRoutes(api, adminService)

	return e.Start(fmt.Sprintf(":%d", cfg.Server.Port))
}
