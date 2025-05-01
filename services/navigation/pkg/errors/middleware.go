package errors

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
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

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, req)
		if err != nil {
			return nil, GRPCErrorHandler(err)
		}
		return resp, nil
	}
}
