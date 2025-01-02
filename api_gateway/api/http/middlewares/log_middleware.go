package middlewares

import (
	"api_gateway/pkg/logging"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type LogMiddleware struct {
	LogService *logging.LogService
}

func (m *LogMiddleware) Handle(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(uint)
	if !ok || userID == 0 {
		return c.JSON(http.StatusUnauthorized)
	}

	companyID, _ := strconv.Atoi(c.Params("company_id"))

	method := c.Method()
	path := c.Path()

	err := m.LogService.LogRequest(userID, uint(companyID), method, path)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to log request",
		})
	}

	return c.Next()
}
