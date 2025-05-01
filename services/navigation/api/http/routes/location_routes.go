package routes

import (
	"github.com/gofiber/fiber/v2"
	"navigation_service/api/handlers"
	"navigation_service/api/http/services"
)

func RegisterLocationRoutes(router fiber.Router, locationService *services.LocationService) {
	handler := handlers.NewLocationHandler(locationService)

	locations := router.Group("/locations")
	locations.Get("/", handler.ListLocations)
	locations.Post("/", handler.CreateLocation)
	locations.Get("/:id", handler.GetLocation)
	locations.Put("/:id", handler.UpdateLocation)
	locations.Delete("/:id", handler.DeleteLocation)
	locations.Get("/type/:type", handler.GetLocationType)
}
