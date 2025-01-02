package http

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"hotels-service/api/http/middleware"
	"hotels-service/api/http/routes"
	"hotels-service/app"
	"hotels-service/config"
	"hotels-service/pkg/errors"
)

func Run(appContainer app.App, cfg config.Config) error {

	NewApp := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			return c.Status(code).JSON(fiber.Map{
				"error":  err.Error(),
				"status": code,
			})
		},
	})

	// Middleware
	NewApp.Use(errors.ErrorMiddleware)
	NewApp.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	api := NewApp.Group("/api/v1")
	routes.RegisterHealthRoutes(NewApp)

	api.Use(middleware.AuthMiddleware())
	routes.RegisterBookingRoutes(appContainer, api)
	routes.RegisterHotelRoutes(appContainer, api)
	routes.RegisterRoomRoutes(appContainer, api)
	routes.RegisterBookingRoutes(appContainer, api)

	return NewApp.Listen(fmt.Sprintf(":%d", cfg.Server.Port))
}
