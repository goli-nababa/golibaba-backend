package handlers

import (
	"github.com/gofiber/fiber/v2"
	"navigation_service/api/http/services"
	"navigation_service/internal/common/types"
	"navigation_service/internal/routing/domain"
	"net/http"
	"strconv"
)

type RoutingHandler struct {
	routingService *services.RoutingService
}

func NewRoutingHandler(routingService *services.RoutingService) *RoutingHandler {
	return &RoutingHandler{
		routingService: routingService,
	}
}

type CreateRouteRequest struct {
	Code         string              `json:"code"`
	FromID       uint                `json:"from_id"`
	ToID         uint                `json:"to_id"`
	VehicleTypes []types.VehicleType `json:"vehicle_types"`
}

type UpdateRouteRequest struct {
	Code         string              `json:"code"`
	FromID       uint                `json:"from_id"`
	ToID         uint                `json:"to_id"`
	VehicleTypes []types.VehicleType `json:"vehicle_types"`
	Active       bool                `json:"active"`
}

type RouteResponse struct {
	ID           uint                `json:"id"`
	UUID         string              `json:"uuid"`
	Code         string              `json:"code"`
	FromID       uint                `json:"from_id"`
	ToID         uint                `json:"to_id"`
	Distance     float64             `json:"distance"`
	VehicleTypes []types.VehicleType `json:"vehicle_types"`
	Active       bool                `json:"active"`
}

func (h *RoutingHandler) CreateRoute(c *fiber.Ctx) error {
	var req CreateRouteRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	route, err := domain.NewRoute(
		req.Code,
		req.FromID,
		req.ToID,
		0,
		req.VehicleTypes,
	)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.routingService.CreateRoute(c.Context(), route); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(RouteResponse{
		ID:           uint(route.ID),
		UUID:         route.UUID,
		Code:         route.Code,
		FromID:       route.FromID,
		ToID:         route.ToID,
		Distance:     route.Distance,
		VehicleTypes: route.VehicleTypes,
		Active:       route.Active,
	})
}

func (h *RoutingHandler) UpdateRoute(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid route ID",
		})
	}

	var req UpdateRouteRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	route := &domain.Routing{
		ID:           domain.RoutingID(uint(id)),
		Code:         req.Code,
		FromID:       req.FromID,
		ToID:         req.ToID,
		VehicleTypes: req.VehicleTypes,
		Active:       req.Active,
	}

	if err := h.routingService.UpdateRoute(c.Context(), route); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(RouteResponse{
		ID:           uint(route.ID),
		UUID:         route.UUID,
		Code:         route.Code,
		FromID:       route.FromID,
		ToID:         route.ToID,
		Distance:     route.Distance,
		VehicleTypes: route.VehicleTypes,
		Active:       route.Active,
	})
}

func (h *RoutingHandler) DeleteRoute(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid route ID",
		})
	}

	if err := h.routingService.DeleteRoute(c.Context(), uint(id)); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(http.StatusNoContent)
}

func (h *RoutingHandler) GetRoute(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid route ID",
		})
	}

	route, err := h.routingService.GetRoute(c.Context(), uint(id))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if route == nil {
		return c.SendStatus(http.StatusNotFound)
	}

	return c.JSON(RouteResponse{
		ID:           uint(route.ID),
		UUID:         route.UUID,
		Code:         route.Code,
		FromID:       route.FromID,
		ToID:         route.ToID,
		Distance:     route.Distance,
		VehicleTypes: route.VehicleTypes,
		Active:       route.Active,
	})
}

func (h *RoutingHandler) GetRouteByUUID(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	route, err := h.routingService.GetRouteByUUID(c.Context(), uuid)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if route == nil {
		return c.SendStatus(http.StatusNotFound)
	}

	return c.JSON(RouteResponse{
		ID:           uint(route.ID),
		UUID:         route.UUID,
		Code:         route.Code,
		FromID:       route.FromID,
		ToID:         route.ToID,
		Distance:     route.Distance,
		VehicleTypes: route.VehicleTypes,
		Active:       route.Active,
	})
}

func (h *RoutingHandler) SearchRoutes(c *fiber.Ctx) error {
	filter := domain.RouteFilter{
		FromID: uint(c.QueryInt("from_id", 0)),
		ToID:   uint(c.QueryInt("to_id", 0)),
		//VehicleType: domain.VehicleType(c.Query("vehicle_type", "")),
		ActiveOnly: c.QueryBool("active_only", false),
	}

	pageSize := c.QueryInt("page_size", 10)
	pageNumber := c.QueryInt("page_number", 1)

	routes, err := h.routingService.FindRoutes(c.Context(), filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := make([]RouteResponse, len(routes))
	for i, route := range routes {
		response[i] = RouteResponse{
			ID:           uint(route.ID),
			UUID:         route.UUID,
			Code:         route.Code,
			FromID:       route.FromID,
			ToID:         route.ToID,
			Distance:     route.Distance,
			VehicleTypes: route.VehicleTypes,
			Active:       route.Active,
		}
	}

	return c.JSON(fiber.Map{
		"data": response,
		"pagination": fiber.Map{
			"page_size":   pageSize,
			"page_number": pageNumber,
			"total":       len(routes),
		},
	})
}
func (h *RoutingHandler) ValidateRoute(c *fiber.Ctx) error {
	routeID := uint(c.QueryInt("route_id"))
	vehicleType := types.VehicleType(c.Query("vehicle_type"))

	if routeID == 0 || vehicleType == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "route_id and vehicle_type are required",
		})
	}

	if err := h.routingService.ValidateRouteForVehicleType(c.Context(), routeID, vehicleType); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"valid": true,
	})
}

func (h *RoutingHandler) ListRoutes(c *fiber.Ctx) error {
	activeOnly := c.QueryBool("active_only", false)
	vehicleType := c.Query("vehicle_type", "")

	filter := domain.RouteFilter{
		ActiveOnly: activeOnly,
	}

	if vehicleType != "" {
		filter.VehicleType = types.VehicleType(vehicleType)
	}

	routes, err := h.routingService.FindRoutes(c.Context(), filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	type RouteResponse struct {
		ID           uint                `json:"id"`
		UUID         string              `json:"uuid"`
		Code         string              `json:"code"`
		FromID       uint                `json:"from_id"`
		ToID         uint                `json:"to_id"`
		Distance     float64             `json:"distance"`
		VehicleTypes []types.VehicleType `json:"vehicle_types"`
		Active       bool                `json:"active"`
		CreatedAt    string              `json:"created_at"`
		UpdatedAt    string              `json:"updated_at"`
	}

	response := make([]RouteResponse, len(routes))
	for i, route := range routes {
		response[i] = RouteResponse{
			ID:           uint(route.ID),
			UUID:         route.UUID,
			Code:         route.Code,
			FromID:       route.FromID,
			ToID:         route.ToID,
			Distance:     route.Distance,
			VehicleTypes: route.VehicleTypes,
			Active:       route.Active,
			CreatedAt:    route.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:    route.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		}
	}

	return c.JSON(response)
}
