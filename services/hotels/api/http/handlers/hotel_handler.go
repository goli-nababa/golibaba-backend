package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"hotels-service/app"
	hotelDomain "hotels-service/internal/hotel/domain"
)

type HotelHandler struct {
	app app.App
}

func NewHotelHandler(app app.App) *HotelHandler {
	return &HotelHandler{app: app}
}

func (h *HotelHandler) CreateHotel(c *fiber.Ctx) error {
	var hotel hotelDomain.Hotel
	if err := c.BodyParser(&hotel); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if err := hotel.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	hotelID, err := h.app.HotelService(c.Context()).CreateHotel(c.Context(), hotel)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"hotel_id": hotelID,
	})
}

func (h *HotelHandler) GetHotel(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid hotel ID")
	}

	hotel, err := h.app.HotelService(c.Context()).GetHotelByID(c.Context(), id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(hotel)
}

func (h *HotelHandler) ListHotels(c *fiber.Ctx) error {
	pageIndex := uint(c.QueryInt("page", 1))
	pageSize := uint(c.QueryInt("size", 10))

	hotels, err := h.app.HotelService(c.Context()).ListHotels(c.Context(), pageIndex, pageSize)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(hotels)
}

func (h *HotelHandler) FindHotels(c *fiber.Ctx) error {
	var filter hotelDomain.HotelFilterItem
	if err := c.QueryParser(&filter); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid filter parameters")
	}

	pageIndex := uint(c.QueryInt("page", 1))
	pageSize := uint(c.QueryInt("size", 10))

	hotels, err := h.app.HotelService(c.Context()).FindHotels(c.Context(), filter, pageIndex, pageSize)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(hotels)
}

func (h *HotelHandler) UpdateHotel(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid hotel ID")
	}

	var hotel hotelDomain.Hotel
	if err := c.BodyParser(&hotel); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if err := hotel.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = h.app.HotelService(c.Context()).UpdateHotel(c.Context(), id, hotel)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *HotelHandler) DeleteHotel(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid hotel ID")
	}

	err = h.app.HotelService(c.Context()).DeleteHotel(c.Context(), id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}
