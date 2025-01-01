package handlers

import (
	"net/http"
	"user_service/api/http/handlers/helpers"
	"user_service/api/http/services"
	"user_service/api/http/types"
	"user_service/app"
	"user_service/config"
	"user_service/internal/domain"

	"github.com/gofiber/fiber/v2"
)

func RegisterDashboardHandlers(router fiber.Router, appContainer app.App, cfg config.ServerConfig) {
	dashboardGroup := router.Group("/dashboard")
	dashboardSvcGetter := services.DashboardServiceGetter(appContainer, cfg)

	dashboardGroup.Get("/notifications", GetNotifications(dashboardSvcGetter))
	dashboardGroup.Post("/notification", CreateNotification(dashboardSvcGetter))
}

func GetNotifications(svcGetter helpers.ServiceGetter[*services.DashboardService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())

		userID, ok := c.Locals("userID").(uint)
		if !ok || userID == 0 {
			return c.JSON(http.StatusUnauthorized)
		}

		notifications, err := svc.GetNotifications(c.Context(), userID)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		response := make([]types.NotificationResponse, len(notifications))
		for i, notif := range notifications {
			response[i] = types.NotificationResponse{
				ID:        uint(notif.ID),
				Message:   notif.Message,
				Seen:      notif.Seen,
				CreatedAt: notif.CreatedAt.String(),
			}
		}

		return c.JSON(response)
	}
}

func CreateNotification(svcGetter helpers.ServiceGetter[*services.DashboardService]) fiber.Handler {
	return func(c *fiber.Ctx) error {
		svc := svcGetter(c.UserContext())

		var req types.CreateNotificationRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		notif := &domain.Notification{
			UserID:  uint(req.UserID),
			Message: req.Message,
			Seen:    false,
		}

		if err := svc.CreateNotification(c.Context(), notif); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(http.StatusCreated).JSON(types.NotificationResponse{
			ID:        uint(notif.ID),
			Message:   notif.Message,
			Seen:      notif.Seen,
			CreatedAt: notif.CreatedAt.String(),
		})
	}
}
