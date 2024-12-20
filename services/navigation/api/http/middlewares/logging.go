package middleware

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"navigation_service/config"
	"navigation_service/internal/common/log"
	"navigation_service/pkg/logging"
	"time"
)

func LoggerMiddleware(cfg *config.Config) fiber.Handler {
	logger := logging.NewLogger(cfg)
	logger.Init()

	return func(c *fiber.Ctx) error {
		startTime := time.Now()

		requestLog := &log.RequestLog{
			Method:    c.Method(),
			Path:      c.Path(),
			IP:        c.IP(),
			StartTime: startTime,
		}

		if userID, ok := c.Locals("userID").(uint); ok {
			requestLog.UserID = userID
		}

		var reqBody interface{}
		if err := c.BodyParser(&reqBody); err == nil {
			requestLog.RequestBody = reqBody
		}

		ctx := context.WithValue(c.UserContext(), log.RequestLogKey, requestLog)
		ctx = context.WithValue(ctx, log.LoggerKey, logger)
		c.SetUserContext(ctx)

		err := c.Next()

		requestLog.StatusCode = c.Response().StatusCode()
		requestLog.ResponseTime = time.Since(startTime)

		if err != nil {
			requestLog.Error = err.Error()
		}

		logData := make(map[logging.ExtraKey]interface{})
		logData[logging.UserID] = requestLog.UserID
		logData[logging.Method] = requestLog.Method
		logData[logging.Path] = requestLog.Path
		logData[logging.StatusCode] = requestLog.StatusCode
		logData[logging.Duration] = requestLog.ResponseTime.String()
		logData[logging.HostIp] = requestLog.IP

		if len(requestLog.HandlerLogs) > 0 {
			logData[logging.ExtraKey("handler_logs")] = requestLog.HandlerLogs
		}

		if requestLog.Error != "" {
			logData[logging.ExtraKey("error")] = requestLog.Error
		}

		switch {
		case requestLog.StatusCode >= 500:
			logger.Error(log.CategoryHTTP, log.SubCategoryAPI, "request_completed", logData)
		case requestLog.StatusCode >= 400:
			logger.Warn(log.CategoryHTTP, log.SubCategoryAPI, "request_completed", logData)
		default:
			logger.Info(log.CategoryHTTP, log.SubCategoryAPI, "request_completed", logData)
		}

		return err
	}
}
