package port

import (
	"context"
	"hotels-service/internal/room/domain"
)

type Service interface {
	CreateRoom(ctx context.Context, room domain.Room) error
	DeleteRoom(ctx context.Context, UUID domain.RoomID) error
	GetAllRooms(ctx context.Context, pageIndex, pageSize uint) ([]domain.Room, error)
	GetAvailableRooms(ctx context.Context, pageIndex, pageSize uint) ([]domain.Room, error)
	FindRoom(ctx context.Context, pageIndex, pageSize uint, filters domain.RoomFilterItem) ([]domain.Room, error)
	GetRoomByID(ctx context.Context, UUID domain.RoomID) (*domain.Room, error)
	SetRoomStatus(ctx context.Context, UUID domain.RoomID, status domain.StatusType) error
	UpdateRoom(ctx context.Context, room domain.Room) error
}
