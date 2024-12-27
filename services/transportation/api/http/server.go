package http

import (
	"fmt"
	"net/http"
	"transportation/app"
	"transportation/config"

	"github.com/labstack/echo"
)

func Run(appContainer app.App, cfg config.ServerConfig) error {
	app := echo.New()

	app.GET("/health", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]bool{"success": true})
	})

	api := app.Group("/api/v1")
	registerCompanyRoutes(api, appContainer, cfg)
	registerTripRoutes(api, appContainer, cfg)
	return app.Start(fmt.Sprintf(":%d", cfg.Port))
}
