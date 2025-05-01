package http

import (
	"fmt"
	"transportation/api/http/services"
	"transportation/api/http/types"

	"github.com/labstack/echo"
)

func CreateCompany(s ServiceGetter[*services.CompanyService]) func(echo.Context) error {
	return func(c echo.Context) error {
		svc := s(c.Request().Context())

		companyRequest := types.CreateCompanyRequest{}

		err := c.Bind(&companyRequest)

		if err != nil {
			fmt.Println(err)
			c.JSON(500, nil)
			return err
		}

		co, err := svc.CreateCompany(c.Request().Context(), companyRequest)
		if err != nil {
			fmt.Println(err)

			c.JSON(500, nil)
			return err
		}

		return c.JSON(200, co)
	}
}

func GetCompanies(s ServiceGetter[*services.CompanyService]) func(echo.Context) error {
	return func(c echo.Context) error {
		svc := s(c.Request().Context())

		companyRequest := types.FilterCompaniesRequest{}

		err := c.Bind(&companyRequest)

		if err != nil {
			c.JSON(500, nil)
			return err
		}

		co, err := svc.GetCompanies(c.Request().Context(), companyRequest)
		if err != nil {
			c.JSON(500, nil)
			return err
		}

		return c.JSON(200, co)
	}
}

//...
