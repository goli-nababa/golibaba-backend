package handlers

import (
	"transportation/api/http/services"

	"github.com/labstack/echo"
)

func CreateCompany(s *services.CompanyService) func(echo.Context) error {
	return func(c echo.Context) error {
		return c.String(200, "test")
	}
}
