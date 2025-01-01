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

func registerTripRoutes(serverGroup *echo.Group, appContainer app.App, cfg config.ServerConfig) {
	s := tripServiceGetter(appContainer, cfg)
	g := serverGroup.Group("/trips")
	g.GET("", GetTrips(s))
	g.GET("/search", SearchTrips(s))
	g.POST("", CreateTrip(s))
	g.POST("/:id/confirm-technical-team", ConfirmTechnicalTeam(s))
	g.POST("/:id/end", EndTrip(s))
	g.POST("/:id/confirm-end", ConfirmEndTrip(s))
	g.POST("/:id/vehicle-requests", CreateVehicleRequest(s))

	//....
}
