package routes

import (
	"github.com/gofiber/fiber/v2"
	"hotels-service/api/http/handlers"
	"hotels-service/app"
)

func RegisterBookingRoutes(appContainer app.App, api fiber.Router) {
	bookingHandler := handlers.NewBookingHandler(appContainer)

	bookings := api.Group("/bookings")
	bookings.Post("/", bookingHandler.CreateBooking)
	bookings.Get("/:id", bookingHandler.GetBooking)
	bookings.Get("/", bookingHandler.GetAllBookings)
	bookings.Get("/search", bookingHandler.FindBookings)
	bookings.Put("/:id", bookingHandler.UpdateBooking)
	bookings.Put("/:id/cancel", bookingHandler.CancelBooking)
	bookings.Delete("/:id", bookingHandler.DeleteBooking)
}
