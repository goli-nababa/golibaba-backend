package routes

import (
	"github.com/gofiber/fiber/v2"
	"hotels-service/api/http/handlers"
	"hotels-service/app"
)

func RegisterRateRoutes(appContainer app.App, api fiber.Router) {
	rateHandler := handlers.NewRateHandler(appContainer)
	rates := api.Group("/rates")
	rates.Post("/", rateHandler.CreateRate)
	rates.Get("/:id", rateHandler.GetRate)
	rates.Get("/", rateHandler.GetAllRates)
	rates.Get("/search", rateHandler.FindRates)
	rates.Put("/:id", rateHandler.UpdateRate)
	rates.Delete("/:id", rateHandler.DeleteRate)
}
