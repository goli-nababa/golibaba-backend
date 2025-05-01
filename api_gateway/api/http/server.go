package http

import (
	"api_gateway/api/http/helpers"
	"api_gateway/api/http/middlewares"
	"api_gateway/api/http/types"
	"api_gateway/config"
	adapters "api_gateway/pkg/adapters/rabbitmq"
	"api_gateway/pkg/logging"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/goli-nababa/golibaba-backend/modules/cache"

	di "api_gateway/app"
)

func Bootstrap(appContainer di.App, cfg config.ServerConfig, grpcCfg config.GrpcConfig) error {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	rabbitMQPublisher := appContainer.RabbitMQPublisher()
	logPublisher := logging.NewRabbitMQLogPublisher(rabbitMQPublisher)
	logService := logging.LogService{Publisher: logPublisher}
	logMiddleware := middlewares.LogMiddleware{LogService: &logService}

	// Add middleware
	app.Use(logMiddleware.Handle)

	// Start RabbitMQ consumer
	appContainer.StartRabbitMQConsumer(func(consumer *adapters.RabbitMQConsumer) {
		if err := consumer.Consume(); err != nil {
			log.Fatalf("Error consuming messages: %v", err)
		}
	})

	app.Post("/v1/register", RegisterServices(appContainer, cfg))

	return app.Listen(fmt.Sprintf(":%d", cfg.Port))
}

func RegisterServices(ac di.App, cfg config.ServerConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := new(types.RegisterRequest)

		if err := helpers.ParseRequestBody(c, body); err != nil {
			fmt.Printf("%v\n", body)
			return c.Status(fiber.StatusBadRequest).JSON(err)
		}

		cacheKey := fmt.Sprintf("%s.%s", body.Name, body.Version)

		cacheProvider := cache.NewJsonObjectCache[*types.RegisterRequest](
			ac.Cache(),
			"service.",
		)

		// Check if the service already exists in the cache
		exists, err := cacheProvider.Exists(c.Context(), cacheKey)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   "failed to check cache existence",
				"details": err.Error(),
			})
		}

		// Save or update the service in the cache
		if exists {
			if err := cacheProvider.Set(c.Context(), cacheKey, time.Duration(cfg.ServiceTTL), body); err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"error":   "failed to update service in cache",
					"details": err.Error(),
				})
			}

			log.Printf("Service %s updated successfuly\n", body.Name)
			log.Printf("Service info: %v\n", body)

			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "service updated successfully",
			})
		}

		if err := cacheProvider.Set(c.Context(), cacheKey, time.Duration(cfg.ServiceTTL), body); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   "failed to save service in cache",
				"details": err.Error(),
			})
		}

		log.Printf("Service %s registered successfuly\n", body.Name)
		log.Printf("Service info: %v\n", body)

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "service registered successfully",
		})
	}
}
