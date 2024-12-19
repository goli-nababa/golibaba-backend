package port

import (
	"context"
	RoomDomain "hotels-service/internal/room/domain"
	UserDomain "hotels-service/internal/user/domain"
)

type Booking interface {
	GetUserByID(ctx context.Context, user UserDomain.UserID) (UserDomain.User, error)
	GetRoomByID(ctx context.Context, room RoomDomain.RoomID) (RoomDomain.Room, error)
}
