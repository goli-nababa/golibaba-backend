package port

import (
	"context"
	"hotels-service/internal/room/domain"
)

type Repo interface {
	Create(ctx context.Context, room domain.Room) (domain.RoomID, error)
	Delete(ctx context.Context, UUID domain.RoomID) error
	Get(ctx context.Context, pageIndex uint, pageSize uint, filter ...domain.RoomFilterItem) ([]domain.Room, error)
	GetByID(ctx context.Context, UUID domain.RoomID) (*domain.Room, error)
	Update(ctx context.Context, UUID domain.RoomID, newData domain.Room) error
}
