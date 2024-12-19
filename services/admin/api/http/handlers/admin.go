package handlers

import (
	"admin/api/http/services"

	"github.com/labstack/echo"
)

func BlockUser(s *services.AdminService) func(echo.Context) error {
	return func(c echo.Context) error {
		return c.String(200, "test")
	}
}
