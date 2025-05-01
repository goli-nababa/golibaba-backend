package handlers

import (
	"github.com/gofiber/fiber/v2"
	"navigation_service/api/http/services"
	"navigation_service/internal/common/types"

	"navigation_service/internal/location/domain"
	"net/http"
	"strconv"
)

type LocationHandler struct {
	locationService *services.LocationService
}

func NewLocationHandler(locationService *services.LocationService) *LocationHandler {
	return &LocationHandler{
		locationService: locationService,
	}
}

type CreateLocationRequest struct {
	Name      string             `json:"name"`
	Type      types.LocationType `json:"type"`
	Address   string             `json:"address"`
	Latitude  float64            `json:"latitude"`
	Longitude float64            `json:"longitude"`
}

type UpdateLocationRequest struct {
	Name      string             `json:"name"`
	Type      types.LocationType `json:"type"`
	Address   string             `json:"address"`
	Latitude  float64            `json:"latitude"`
	Longitude float64            `json:"longitude"`
	Active    bool               `json:"active"`
}

type LocationResponse struct {
	ID        uint               `json:"id"`
	Name      string             `json:"name"`
	Type      types.LocationType `json:"type"`
	Address   string             `json:"address"`
	Latitude  float64            `json:"latitude"`
	Longitude float64            `json:"longitude"`
	Active    bool               `json:"active"`
}

func (h *LocationHandler) CreateLocation(c *fiber.Ctx) error {
	var req CreateLocationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	location := &domain.Location{
		Name:      req.Name,
		Type:      req.Type,
		Address:   req.Address,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Active:    true,
	}

	if err := h.locationService.CreateLocation(c.Context(), location); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(LocationResponse{
		ID:        uint(location.ID),
		Name:      location.Name,
		Type:      location.Type,
		Address:   location.Address,
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
		Active:    location.Active,
	})
}

func (h *LocationHandler) UpdateLocation(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid location ID",
		})
	}

	var req UpdateLocationRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	location := &domain.Location{
		ID:        domain.LocationID(uint(id)),
		Name:      req.Name,
		Type:      req.Type,
		Address:   req.Address,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Active:    req.Active,
	}

	if err := h.locationService.UpdateLocation(c.Context(), location); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(LocationResponse{
		ID:        uint(location.ID),
		Name:      location.Name,
		Type:      location.Type,
		Address:   location.Address,
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
		Active:    location.Active,
	})
}

func (h *LocationHandler) DeleteLocation(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid location ID",
		})
	}

	if err := h.locationService.DeleteLocation(c.Context(), uint(id)); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(http.StatusNoContent)
}

func (h *LocationHandler) GetLocation(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid location ID",
		})
	}

	location, err := h.locationService.GetLocation(c.Context(), uint(id))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if location == nil {
		return c.SendStatus(http.StatusNotFound)
	}

	return c.JSON(LocationResponse{
		ID:        uint(location.ID),
		Name:      location.Name,
		Type:      location.Type,
		Address:   location.Address,
		Latitude:  location.Latitude,
		Longitude: location.Longitude,
		Active:    location.Active,
	})
}

func (h *LocationHandler) ListLocations(c *fiber.Ctx) error {
	activeOnly := c.QueryBool("active_only", false)

	locations, err := h.locationService.ListLocations(c.Context(), activeOnly)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := make([]LocationResponse, len(locations))
	for i, loc := range locations {
		response[i] = LocationResponse{
			ID:        uint(loc.ID),
			Name:      loc.Name,
			Type:      loc.Type,
			Address:   loc.Address,
			Latitude:  loc.Latitude,
			Longitude: loc.Longitude,
			Active:    loc.Active,
		}
	}

	return c.JSON(response)
}

func (h *LocationHandler) GetLocationType(c *fiber.Ctx) error {
	locationType := types.LocationType(c.Params("type"))

	locations, err := h.locationService.GetLocationsByType(c.Context(), locationType)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := make([]LocationResponse, len(locations))
	for i, loc := range locations {
		response[i] = LocationResponse{
			ID:        uint(loc.ID),
			Name:      loc.Name,
			Type:      loc.Type,
			Address:   loc.Address,
			Latitude:  loc.Latitude,
			Longitude: loc.Longitude,
			Active:    loc.Active,
		}
	}

	return c.JSON(response)
}
