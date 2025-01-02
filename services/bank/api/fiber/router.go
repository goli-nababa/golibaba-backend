package fiber

import (
	"bank_service/api/fiber/handlers"
	"bank_service/app"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Router struct {
	app app.App
}

func NewRouter(app app.App) *Router {
	return &Router{app: app}
}

func (r *Router) SetupRoutes(app *fiber.App) {

	app.Use(cors.New())

	paymentHandler := handlers.NewPaymentHandler(r.app)

	api := app.Group("/api/v1")

	payments := api.Group("/payments")

	payments.Get("/callback", paymentHandler.HandleZarinpalCallback)

}
