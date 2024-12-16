package handlers

import (
	"github.com/gofiber/fiber/v2"
	"user_service/app"
)

func RegisterAccountHandlers(router fiber.Router, appContainer app.App, cfg config.ServerConfig) {
	//accountGroup := router.Group("/account")
	//accountSvcGetter := services.AccountServiceGetter(appContainer, cfg)
}
