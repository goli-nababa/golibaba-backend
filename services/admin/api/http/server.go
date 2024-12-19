package http

import (
	"admin/api/http/routes"
	"admin/app"
	"admin/config"
	"context"
	"fmt"

	"github.com/labstack/echo"
)

func Run(appContainer app.App, cfg config.ServerConfig) error {
	app := echo.New()
	api := app.Group("/api/v1")
	ctx := context.Background()
	routes.RegisterAdminRoutes(ctx, api, appContainer, cfg)

	return app.Start(fmt.Sprintf(":%d", cfg.Port))
}
