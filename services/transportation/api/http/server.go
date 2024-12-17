package http

import (
	"context"
	"fmt"
	"transportation/api/http/routes"
	"transportation/app"
	"transportation/config"

	"github.com/labstack/echo"
)

func Run(appContainer app.App, cfg config.ServerConfig) error {
	app := echo.New()
	api := app.Group("/api/v1")
	ctx := context.Background()
	routes.RegisterCompanyRoutes(ctx, api, appContainer, cfg)

	return app.Start(fmt.Sprintf(":%d", cfg.Port))
}
