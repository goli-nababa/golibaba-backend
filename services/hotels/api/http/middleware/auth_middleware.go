package middleware

import "github.com/gofiber/fiber/v2"

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Get("X-User-ID")
		if userID == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "User not authenticated")
		}

		return c.Next()
	}
}
