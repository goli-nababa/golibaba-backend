package http

import (
	"api_gateway/config"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"

	di "api_gateway/app"
)

func Bootstrap(appContainer di.App, cfg config.ServerConfig) error {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Post("/register", RegisterServices(appContainer, cfg))

	return app.Listen(fmt.Sprintf(":%d", cfg.Port))
}

func RegisterServices(appContainer di.App, cfg config.ServerConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	}
}
