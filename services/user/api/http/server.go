package http

import (
	"encoding/json"
	"fmt"
	"user_service/api/http/handlers"
	di "user_service/app"
	"user_service/config"

	"github.com/gofiber/fiber/v2"
)

func Bootstrap(appContainer di.App, cfg config.ServerConfig) error {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	api := app.Group("/api/v1")

	handlers.RegisterAccountHandlers(api, appContainer, cfg)
	handlers.RegisterDashboardHandlers(api, appContainer, cfg)

	return app.Listen(fmt.Sprintf(":%d", cfg.Port))
}
