package room

import (
	"context"
	hotelPort "hotels-service/internal/hotel/port"
	ratePort "hotels-service/internal/rate/port"
	"hotels-service/internal/room/domain"
	"hotels-service/internal/room/port"
)

type service struct {
	reateService ratePort.Service
	hotelService hotelPort.Service
	repo         port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{repo: repo}
}

func (s *service) CreateRoom(ctx context.Context, room domain.Room) error {
	_, err := s.repo.Create(ctx, room)
	return err
}

func (s *service) UpdateRoom(ctx context.Context, room domain.Room) error {
	return s.repo.Update(ctx, room.ID, room)
}

func (s *service) DeleteRoom(ctx context.Context, UUID domain.RoomID) error {
	return s.repo.Delete(ctx, UUID)
}

func (s *service) FindRoom(ctx context.Context, pageIndex, pageSize uint, filters domain.RoomFilterItem) ([]domain.Room, error) {
	return s.repo.Get(ctx, pageIndex, pageSize, filters)
}

func (s *service) GetRoomByID(ctx context.Context, UUID domain.RoomID) (*domain.Room, error) {
	return s.repo.GetByID(ctx, UUID)
}

func (s *service) GetAllRooms(ctx context.Context, pageIndex, pageSize uint) ([]domain.Room, error) {
	return s.repo.Get(ctx, pageIndex, pageSize)
}

func (s *service) GetAvailableRooms(ctx context.Context, pageIndex, pageSize uint) ([]domain.Room, error) {
	filter := domain.RoomFilterItem{
		Status: domain.StatusTypeAvailable,
	}
	return s.repo.Get(ctx, pageIndex, pageSize, filter)
}

func (s *service) SetRoomStatus(ctx context.Context, UUID domain.RoomID, status domain.StatusType) error {
	room, err := s.repo.GetByID(ctx, UUID)
	if err != nil {
		return err
	}

	room.Status = status
	return s.repo.Update(ctx, UUID, *room)
}
