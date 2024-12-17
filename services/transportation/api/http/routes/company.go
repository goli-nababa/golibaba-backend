package routes

import (
	"context"
	"transportation/api/http/handlers"
	"transportation/api/http/services"
	"transportation/app"
	"transportation/config"

	"github.com/labstack/echo"
)

func RegisterCompanyRoutes(ctx context.Context, serverGroup *echo.Group, appContainer app.App, cfg config.ServerConfig) {
	s := services.NewCompanyService(appContainer.CompanyService(ctx))
	g := serverGroup.Group("/companies")
	g.POST("", handlers.CreateCompany(s))
}
