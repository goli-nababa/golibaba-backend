package routes

import (
	"github.com/gofiber/fiber/v2"
	"hotels-service/api/http/handlers"
	"hotels-service/app"
)

func RegisterHotelRoutes(appContainer app.App, api fiber.Router) {
	hotelHandler := handlers.NewHotelHandler(appContainer)
	hotels := api.Group("/hotels")
	hotels.Post("/", hotelHandler.CreateHotel)
	hotels.Get("/:id", hotelHandler.GetHotel)
	hotels.Get("/", hotelHandler.ListHotels)
	hotels.Get("/search", hotelHandler.FindHotels)
	hotels.Put("/:id", hotelHandler.UpdateHotel)
	hotels.Delete("/:id", hotelHandler.DeleteHotel)
}
