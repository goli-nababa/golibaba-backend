package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goli-nababa/golibaba-backend/common"
	"github.com/google/uuid"
	"hotels-service/app"
	"hotels-service/internal/booking/domain"
	"strconv"
)

type BookingHandler struct {
	app app.App
}

func NewBookingHandler(app app.App) *BookingHandler {
	return &BookingHandler{app: app}
}

func (h *BookingHandler) CreateBooking(c *fiber.Ctx) error {
	// Get user ID from gateway header
	userIDStr := c.Get("X-User-ID")
	if userIDStr == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "User not authenticated")
	}

	id, err := strconv.Atoi(userIDStr)
	if err != nil {
		return err
	}
	userID := common.UserID(id)

	// Check if user has permission to create booking
	hasAccess, err := h.app.UserServiceClient().CheckAccess(userID, []string{"create:booking"})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to check permissions")
	}
	if !hasAccess {
		return fiber.NewError(fiber.StatusForbidden, "Insufficient permissions")
	}

	var booking domain.Booking
	if err := c.BodyParser(&booking); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	// Set user ID from authenticated user
	booking.UserID = uuid.MustParse(userIDStr)

	if err := booking.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	bookingID, err := h.app.BookingService(c.Context()).CreateNewBooking(c.Context(), booking)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"booking_id": bookingID,
	})
}

func (h *BookingHandler) GetBooking(c *fiber.Ctx) error {
	userIDStr := c.Get("X-User-ID")
	if userIDStr == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "User not authenticated")
	}

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid booking ID")
	}

	// Get the booking
	booking, err := h.app.BookingService(c.Context()).GetBookingByID(c.Context(), id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	_id, err := strconv.Atoi(userIDStr)
	if err != nil {
		return err
	}
	userID := common.UserID(_id)

	if booking.UserID.String() != userIDStr {
		hasAccess, err := h.app.UserServiceClient().CheckAccess(userID, []string{"view:any_booking"})
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to check permissions")
		}
		if !hasAccess {
			return fiber.NewError(fiber.StatusForbidden, "Insufficient permissions")
		}
	}

	return c.JSON(booking)
}

func (h *BookingHandler) GetAllBookings(c *fiber.Ctx) error {
	pageIndex := uint(c.QueryInt("page", 1))
	pageSize := uint(c.QueryInt("size", 10))

	bookings, err := h.app.BookingService(c.Context()).GetAllBooking(c.Context(), pageIndex, pageSize)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(bookings)
}

func (h *BookingHandler) FindBookings(c *fiber.Ctx) error {
	var filter domain.BookingFilterItem
	if err := c.QueryParser(&filter); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid filter parameters")
	}

	pageIndex := uint(c.QueryInt("page", 1))
	pageSize := uint(c.QueryInt("size", 10))

	bookings, err := h.app.BookingService(c.Context()).FindBooking(c.Context(), filter, pageIndex, pageSize)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(bookings)
}

func (h *BookingHandler) UpdateBooking(c *fiber.Ctx) error {
	userIDStr := c.Get("X-User-ID")
	if userIDStr == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "User not authenticated")
	}

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid booking ID")
	}

	bookingData, err := h.app.BookingService(c.Context()).GetBookingByID(c.Context(), id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	_id, err := strconv.Atoi(userIDStr)
	if err != nil {
		return err
	}
	userID := common.UserID(_id)

	if bookingData.UserID.String() != userIDStr {
		hasAccess, err := h.app.UserServiceClient().CheckAccess(userID, []string{"update:any_booking"})
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to check permissions")
		}
		if !hasAccess {
			return fiber.NewError(fiber.StatusForbidden, "Insufficient permissions")
		}
	}

	var booking domain.Booking
	if err := c.BodyParser(&booking); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if err := booking.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = h.app.BookingService(c.Context()).EditeBooking(c.Context(), id, booking)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *BookingHandler) CancelBooking(c *fiber.Ctx) error {
	userIDStr := c.Get("X-User-ID")
	if userIDStr == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "User not authenticated")
	}

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid booking ID")
	}

	booking, err := h.app.BookingService(c.Context()).GetBookingByID(c.Context(), id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	_id, err := strconv.Atoi(userIDStr)
	if err != nil {
		return err
	}
	userID := common.UserID(_id)

	if booking.UserID.String() != userIDStr {
		hasAccess, err := h.app.UserServiceClient().CheckAccess(userID, []string{"cancel:any_booking"})
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to check permissions")
		}
		if !hasAccess {
			return fiber.NewError(fiber.StatusForbidden, "Insufficient permissions")
		}
	}

	err = h.app.BookingService(c.Context()).CancelBooking(c.Context(), id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *BookingHandler) DeleteBooking(c *fiber.Ctx) error {
	userIDStr := c.Get("X-User-ID")
	if userIDStr == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "User not authenticated")
	}
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid booking ID")
	}

	_id, err := strconv.Atoi(userIDStr)
	if err != nil {
		return err
	}
	userID := common.UserID(_id)

	booking, err := h.app.BookingService(c.Context()).GetBookingByID(c.Context(), id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	if booking.UserID.String() != userIDStr {
		hasAccess, err := h.app.UserServiceClient().CheckAccess(userID, []string{"delete:any_booking"})
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to check permissions")
		}
		if !hasAccess {
			return fiber.NewError(fiber.StatusForbidden, "Insufficient permissions")
		}
	}

	err = h.app.BookingService(c.Context()).DeleteBooking(c.Context(), id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}
