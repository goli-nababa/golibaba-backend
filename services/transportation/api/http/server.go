package http

import (
	"fmt"
	"transportation/app"
	"transportation/config"

	"github.com/labstack/echo"
)

func Run(appContainer app.App, cfg config.ServerConfig) error {
	app := echo.New()
	api := app.Group("/api/v1")
	registerCompanyRoutes(api, appContainer, cfg)
	registerTripRoutes(api, appContainer, cfg)
	return app.Start(fmt.Sprintf(":%d", cfg.Port))
}
