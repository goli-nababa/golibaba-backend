package http

import (
	"transportation/app"
	"transportation/config"

	"github.com/labstack/echo"
)

func registerCompanyRoutes(serverGroup *echo.Group, appContainer app.App, cfg config.ServerConfig) {
	s := companyServiceGetter(appContainer, cfg)
	g := serverGroup.Group("/companies")
	g.POST("", CreateCompany(s))
	g.GET("", GetCompanies(s))

	//....
}
