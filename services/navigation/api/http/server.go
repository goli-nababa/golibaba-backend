package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"navigation_service/api/http/middlewares"
	"navigation_service/api/http/routes"
	"navigation_service/api/http/services"
	di "navigation_service/app"
	"navigation_service/config"
)

func Bootstrap(appContainer di.App, cfg config.Config) error {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Static files
	app.Static("/", "./web")

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(middleware.LoggerMiddleware(&cfg))
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Initialize services
	locationService := services.NewLocationService(appContainer.LocationService(context.Background()))
	routingService := services.NewRoutingService(appContainer.RoutingService(context.Background()))

	// API Routes
	api := app.Group("/api/v1")

	// Register route groups
	routes.RegisterHealthRoutes(app)
	routes.RegisterLocationRoutes(api, locationService)
	routes.RegisterRouteRoutes(api, routingService)

	return app.Listen(fmt.Sprintf(":%d", cfg.Server.Port))
}
