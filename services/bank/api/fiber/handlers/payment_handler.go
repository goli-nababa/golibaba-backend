package handlers

import (
	"bank_service/app"
	domain2 "bank_service/internal/services/payment/domain"
	"github.com/gofiber/fiber/v2"
)

type PaymentHandler struct {
	app app.App
}

func NewPaymentHandler(app app.App) *PaymentHandler {
	return &PaymentHandler{app: app}
}

func (r *PaymentHandler) HandleZarinpalCallback(c *fiber.Ctx) error {
	authority := c.Query("Authority")
	status := c.Query("Status")
	if authority == "" || status == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "missing required parameters",
		})
	}

	callback := &domain2.PaymentCallback{
		Authority: authority,
		Status:    status,
	}

	_, redirectURL, err := r.app.PaymentService(c.Context()).HandlePaymentCallback(c.Context(), callback)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{

		"message": redirectURL,
	})
}
