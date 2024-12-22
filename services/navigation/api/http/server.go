package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"navigation_service/api/handlers"
	middleware "navigation_service/api/http/middlewares"
	"navigation_service/api/http/services"
	di "navigation_service/app"
	"navigation_service/config"
	"time"
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

	// Services
	locationService := services.NewLocationService(appContainer.LocationService(context.Background()))
	routingService := services.NewRoutingService(appContainer.RoutingService(context.Background()))

	// Handlers
	locationHandler := handlers.NewLocationHandler(locationService)
	routingHandler := handlers.NewRoutingHandler(routingService)

	// API Routes
	api := app.Group("/api/v1")

	// Location Routes
	locations := api.Group("/locations")
	locations.Get("/", locationHandler.ListLocations)

	locations.Post("/", locationHandler.CreateLocation)
	locations.Get("/:id", locationHandler.GetLocation)
	locations.Put("/:id", locationHandler.UpdateLocation)
	locations.Delete("/:id", locationHandler.DeleteLocation)
	locations.Get("/type/:type", locationHandler.GetLocationType) // Fixed method name

	// Routing Routes
	routes := api.Group("/routes")
	routes.Get("/search", routingHandler.SearchRoutes)
	routes.Get("/", routingHandler.ListRoutes)

	routes.Post("/", routingHandler.CreateRoute)
	routes.Get("/:id", routingHandler.GetRoute)
	routes.Get("/uuid/:uuid", routingHandler.GetRouteByUUID)
	routes.Put("/:id", routingHandler.UpdateRoute)
	routes.Delete("/:id", routingHandler.DeleteRoute)
	routes.Post("/validate", routingHandler.ValidateRoute)

	// Health Check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
			"time":   time.Now().Format(time.RFC3339),
		})
	})

	return app.Listen(fmt.Sprintf(":%d", cfg.Server.Port))
}
