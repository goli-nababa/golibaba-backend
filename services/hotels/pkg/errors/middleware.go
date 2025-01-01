package errors

import (
	"github.com/gofiber/fiber/v2"
)

func ErrorMiddleware(c *fiber.Ctx) error {
	defer func() {
		if err := recover(); err != nil {
			customErr := NewError(ErrInternal, "خطای غیرمنتظره", fiber.StatusInternalServerError)
			_ = c.Status(customErr.StatusCode).JSON(fiber.Map{
				"success": false,
				"error": fiber.Map{
					"code":    customErr.Code,
					"message": customErr.Message,
					"details": customErr.Details,
				},
			})
		}
	}()
	return c.Next()
}
