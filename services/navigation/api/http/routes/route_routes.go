package routes

import (
	"github.com/gofiber/fiber/v2"
	"navigation_service/api/handlers"
	"navigation_service/api/http/services"
)

func RegisterRouteRoutes(router fiber.Router, routingService *services.RoutingService) {
	handler := handlers.NewRoutingHandler(routingService)

	routes := router.Group("/routes")
	routes.Get("/search", handler.SearchRoutes)
	routes.Get("/", handler.ListRoutes)
	routes.Post("/", handler.CreateRoute)
	routes.Get("/:id", handler.GetRoute)
	routes.Get("/uuid/:uuid", handler.GetRouteByUUID)
	routes.Put("/:id", handler.UpdateRoute)
	routes.Delete("/:id", handler.DeleteRoute)
	routes.Post("/validate", handler.ValidateRoute)
}
