package http

import (
	"strconv"
	"transportation/api/http/services"
	"transportation/internal/trip/domain"

	"github.com/labstack/echo"
)

func CreateTrip(s ServiceGetter[*services.TripService]) func(echo.Context) error {
	return func(c echo.Context) error {
		svc := s(c.Request().Context())

		tripRequest := domain.CreateTripRequest{}

		err := c.Bind(&tripRequest)

		if err != nil {
			c.JSON(400, map[string]string{"message": err.Error()})
			return err
		}

		t, err := svc.CreateTrip(c.Request().Context(), tripRequest)

		if err != nil {

			c.JSON(500, map[string]string{"message": err.Error()})
			return err
		}

		return c.JSON(200, t)
	}
}

func GetTrips(s ServiceGetter[*services.TripService]) func(echo.Context) error {
	return func(c echo.Context) error {
		svc := s(c.Request().Context())

		tripRequest := domain.GetTripsRequest{}

		err := c.Bind(&tripRequest)

		if err != nil {
			c.JSON(400, nil)
			return err
		}

		t, err := svc.MainService.GetTrips(c.Request().Context(), tripRequest)
		if err != nil {

			c.JSON(500, map[string]string{"message": err.Error()})
			return err
		}

		return c.JSON(200, t)
	}
}

func SearchTrips(s ServiceGetter[*services.TripService]) func(echo.Context) error {
	return func(c echo.Context) error {
		svc := s(c.Request().Context())

		tripRequest := domain.GetTripsRequest{}

		err := c.Bind(&tripRequest)

		if err != nil {
			c.JSON(400, map[string]string{"message": err.Error()})
			return err
		}

		t, err := svc.MainService.SearchTrips(c.Request().Context(), tripRequest)
		if err != nil {

			c.JSON(500, map[string]string{"message": err.Error()})
			return err
		}

		return c.JSON(200, t)
	}
}

func ConfirmTechnicalTeam(s ServiceGetter[*services.TripService]) func(echo.Context) error {
	return func(c echo.Context) error {
		svc := s(c.Request().Context())

		id := c.Param("id")

		iid, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(400, map[string]string{"message": "id is not valid"})
			return err
		}

		t, err := svc.ConfirmTechnicalTeam(c.Request().Context(), domain.TripId(iid))
		if err != nil {

			c.JSON(500, map[string]string{"message": err.Error()})
			return err
		}

		return c.JSON(200, t)
	}
}

func EndTrip(s ServiceGetter[*services.TripService]) func(echo.Context) error {
	return func(c echo.Context) error {
		svc := s(c.Request().Context())

		id := c.Param("id")

		iid, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(400, map[string]string{"message": "id is not valid"})
			return err
		}

		t, err := svc.EndTrip(c.Request().Context(), domain.TripId(iid))
		if err != nil {

			c.JSON(500, map[string]string{"message": err.Error()})
			return err
		}

		return c.JSON(200, t)
	}
}

func ConfirmEndTrip(s ServiceGetter[*services.TripService]) func(echo.Context) error {
	return func(c echo.Context) error {
		svc := s(c.Request().Context())

		id := c.Param("id")

		iid, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(400, map[string]string{"message": "id is not valid"})
			return err
		}

		t, err := svc.MainService.ConfirmEndTrip(c.Request().Context(), domain.TripId(iid))
		if err != nil {

			c.JSON(500, map[string]string{"message": err.Error()})
			return err
		}

		return c.JSON(200, t)
	}
}

func CreateVehicleRequest(s ServiceGetter[*services.TripService]) func(echo.Context) error {
	return func(c echo.Context) error {
		svc := s(c.Request().Context())

		req := domain.CreateVehicleRequest{}

		err := c.Bind(&req)

		if err != nil {
			c.JSON(400, err.Error())
			return err
		}

		id := c.Param("id")

		iid, err := strconv.Atoi(id)

		if err != nil {
			c.JSON(400, map[string]string{"message": "id is not valid"})
			return err
		}

		req.TripId = domain.TripId(iid)

		t, err := svc.MainService.CreateVehicleRequest(c.Request().Context(), req)
		if err != nil {

			c.JSON(500, map[string]string{"message": err.Error()})
			return err
		}

		return c.JSON(200, t)
	}
}
