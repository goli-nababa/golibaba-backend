package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goli-nababa/golibaba-backend/common"
	"github.com/google/uuid"
	"hotels-service/app"
	"hotels-service/internal/room/domain"
	"strconv"
)

type RoomHandler struct {
	app app.App
}

func NewRoomHandler(app app.App) *RoomHandler {
	return &RoomHandler{app: app}
}

func (h *RoomHandler) CreateRoom(c *fiber.Ctx) error {
	userIDStr := c.Get("X-User-ID")
	if userIDStr == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "User not authenticated")
	}

	id, err := strconv.Atoi(userIDStr)
	if err != nil {
		return err
	}
	userID := common.UserID(id)

	// Check if user has permission to create rooms
	hasAccess, err := h.app.UserServiceClient().CheckAccess(userID, []string{"create:room"})
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to check permissions")
	}
	if !hasAccess {
		return fiber.NewError(fiber.StatusForbidden, "Insufficient permissions")
	}

	var room domain.Room
	if err := c.BodyParser(&room); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	// Get hotel to verify ownership
	hotel, err := h.app.HotelService(c.Context()).GetHotelByID(c.Context(), room.HotelID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Hotel not found")
	}

	// Verify user owns the hotel or has admin permissions
	if hotel.OwnerID.String() != userIDStr {
		hasAccess, err := h.app.UserServiceClient().CheckAccess(userID, []string{"create:any_room"})
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to check permissions")
		}
		if !hasAccess {
			return fiber.NewError(fiber.StatusForbidden, "Insufficient permissions")
		}
	}

	if err := room.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = h.app.RoomService(c.Context()).CreateRoom(c.Context(), room)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *RoomHandler) GetRoom(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid room ID")
	}

	room, err := h.app.RoomService(c.Context()).GetRoomByID(c.Context(), id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(room)
}

func (h *RoomHandler) GetAllRooms(c *fiber.Ctx) error {
	pageIndex := uint(c.QueryInt("page", 1))
	pageSize := uint(c.QueryInt("size", 10))

	rooms, err := h.app.RoomService(c.Context()).GetAllRooms(c.Context(), pageIndex, pageSize)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(rooms)
}

func (h *RoomHandler) GetAvailableRooms(c *fiber.Ctx) error {
	pageIndex := uint(c.QueryInt("page", 1))
	pageSize := uint(c.QueryInt("size", 10))

	rooms, err := h.app.RoomService(c.Context()).GetAvailableRooms(c.Context(), pageIndex, pageSize)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(rooms)
}

func (h *RoomHandler) FindRooms(c *fiber.Ctx) error {
	var filter domain.RoomFilterItem
	if err := c.QueryParser(&filter); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid filter parameters")
	}

	pageIndex := uint(c.QueryInt("page", 1))
	pageSize := uint(c.QueryInt("size", 10))

	rooms, err := h.app.RoomService(c.Context()).FindRoom(c.Context(), pageIndex, pageSize, filter)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.JSON(rooms)
}

func (h *RoomHandler) UpdateRoom(c *fiber.Ctx) error {
	userIDStr := c.Get("X-User-ID")
	if userIDStr == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "User not authenticated")
	}

	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid room ID")
	}

	existingRoom, err := h.app.RoomService(c.Context()).GetRoomByID(c.Context(), id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	hotel, err := h.app.HotelService(c.Context()).GetHotelByID(c.Context(), existingRoom.HotelID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Hotel not found")
	}

	_id, err := strconv.Atoi(userIDStr)
	if err != nil {
		return err
	}
	userID := common.UserID(_id)

	if hotel.OwnerID.String() != userIDStr {
		hasAccess, err := h.app.UserServiceClient().CheckAccess(userID, []string{"update:any_room"})
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to check permissions")
		}
		if !hasAccess {
			return fiber.NewError(fiber.StatusForbidden, "Insufficient permissions")
		}
	}

	var room domain.Room
	if err := c.BodyParser(&room); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	room.ID = id

	if err := room.Validate(); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = h.app.RoomService(c.Context()).UpdateRoom(c.Context(), room)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}
func (h *RoomHandler) DeleteRoom(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid room ID")
	}

	err = h.app.RoomService(c.Context()).DeleteRoom(c.Context(), id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *RoomHandler) SetRoomStatus(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid room ID")
	}

	var req struct {
		Status domain.StatusType `json:"status"`
	}

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}

	err = h.app.RoomService(c.Context()).SetRoomStatus(c.Context(), id, req.Status)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.SendStatus(fiber.StatusOK)
}
