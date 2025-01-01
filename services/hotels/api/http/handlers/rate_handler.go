package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goli-nababa/golibaba-backend/common"
	"github.com/google/uuid"
	"hotels-service/app"
	"hotels-service/internal/rate/domain"
	"strconv"
)

type RateHandler struct {
	app app.App
}

func NewRateHandler(app app.App) *RateHandler {
	return &RateHandler{app: app}
}

func (h *RateHandler) CreateRate(c *fiber.Ctx) error {
	userIDStr := c.Get("X-User-ID")
	if userIDStr == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "User not authenticated")
	}

	_id, err := strconv.Atoi(userIDStr)
	if err != nil {
		return err
	}
	userID := common.UserID(_id)

	// Check if user has permission to create rates
	hasAccess, err := h.app.UserServiceClient().CheckAccess(userID, []string{"create:rate"})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to check permissions")
	}
	if !hasAccess {
		return fiber.NewError(fiber.StatusForbidden, "Insufficient permissions")
	}

	var rate domain.Rate
	if err := c.BodyParser(&rate); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if err := rate.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	rateID, err := h.app.RateService(c.Context()).CreateNewRate(c.Context(), rate)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"rate_id": rateID,
	})
}

func (h *RateHandler) GetRate(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid rate ID")
	}

	rate, err := h.app.RateService(c.Context()).GetRateByID(c.Context(), id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(rate)
}

func (h *RateHandler) GetAllRates(c *fiber.Ctx) error {
	pageIndex := uint(c.QueryInt("page", 1))
	pageSize := uint(c.QueryInt("size", 10))

	rates, err := h.app.RateService(c.Context()).GetAllRate(c.Context(), pageIndex, pageSize)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(rates)
}

func (h *RateHandler) FindRates(c *fiber.Ctx) error {
	var filter domain.RateFilterItem
	if err := c.QueryParser(&filter); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid filter parameters")
	}

	pageIndex := uint(c.QueryInt("page", 1))
	pageSize := uint(c.QueryInt("size", 10))

	rates, err := h.app.RateService(c.Context()).FindRate(c.Context(), filter, pageIndex, pageSize)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(rates)
}

func (h *RateHandler) UpdateRate(c *fiber.Ctx) error {
	userIDStr := c.Get("X-User-ID")
	if userIDStr == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "User not authenticated")
	}

	_id, err := strconv.Atoi(userIDStr)
	if err != nil {
		return err
	}
	userID := common.UserID(_id)

	// Check if user has permission to update rates
	hasAccess, err := h.app.UserServiceClient().CheckAccess(userID, []string{"update:rate"})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to check permissions")
	}
	if !hasAccess {
		return fiber.NewError(fiber.StatusForbidden, "Insufficient permissions")
	}

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid rate ID")
	}

	var rate domain.Rate
	if err := c.BodyParser(&rate); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	if err := rate.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = h.app.RateService(c.Context()).EditeRate(c.Context(), id, rate)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *RateHandler) DeleteRate(c *fiber.Ctx) error {
	userIDStr := c.Get("X-User-ID")
	if userIDStr == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "User not authenticated")
	}

	_id, err := strconv.Atoi(userIDStr)
	if err != nil {
		return err
	}
	userID := common.UserID(_id)

	// Check if user has permission to delete rates
	hasAccess, err := h.app.UserServiceClient().CheckAccess(userID, []string{"delete:rate"})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to check permissions")
	}
	if !hasAccess {
		return fiber.NewError(fiber.StatusForbidden, "Insufficient permissions")
	}

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid rate ID")
	}

	err = h.app.RateService(c.Context()).DeleteRate(c.Context(), id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}
