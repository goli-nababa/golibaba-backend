package http

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	di "user_service/app"
	"user_service/config"
)

func Bootstrap(appContainer di.App, cfg config.ServerConfig) error {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	/*api := app.Group("/api/v1", middlerwares.SetUserContext)

	handlers.RegisterAccountHandlers(api, appContainer, cfg)*/

	return app.Listen(fmt.Sprintf(":%d", cfg.Port))
}
