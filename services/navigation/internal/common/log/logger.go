package log

import "github.com/gofiber/fiber/v2"

func AddLogEntry(c *fiber.Ctx, message string, data map[string]interface{}) {
	ctx := c.UserContext()
	if requestLog, ok := ctx.Value(RequestLogKey).(*RequestLog); ok {
		requestLog.HandlerLogs = append(requestLog.HandlerLogs, HandlerLog{
			Message: message,
			Data:    data,
		})
	}
}
