package port

import (
	"context"
	roomDomain "hotels-service/internal/room/domain"
	userDomain "hotels-service/internal/user/domain"
)

type Repo interface {
	GetUserByID(ctx context.Context, user userDomain.UserID) (userDomain.User, error)
	GetRoomByID(ctx context.Context, room roomDomain.RoomID) (roomDomain.Room, error)
}
