package storage

import (
	"context"
	"hotels-service/internal/room/domain"
	roomPort "hotels-service/internal/room/port"
	"hotels-service/pkg/adapters/storage/types"
	"time"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type roomRepo struct {
	db *gorm.DB
}

func NewRoomRepo(db *gorm.DB) roomPort.Repo {
	return &roomRepo{
		db: db,
	}
}

func (r *roomRepo) Create(ctx context.Context, newRecord domain.Room) (domain.RoomID, error) {
	room := new(types.Room)
	copier.Copy(room, &newRecord)
	room.CreateAt = time.Now()
	room.UpdateAt = time.Now()
	result := r.db.WithContext(ctx).Create(room)
	return newRecord.ID, result.Error
}
func (r *roomRepo) GetByID(ctx context.Context, UUID domain.RoomID) (*domain.Room, error) {
	roomDomain := new(domain.Room)
	room := new(types.Room)
	result := r.db.WithContext(ctx).First(room, UUID)
	copier.Copy(roomDomain, room)
	return roomDomain, result.Error
}
func (r *roomRepo) Get(ctx context.Context, pageIndex, pageSize uint, filter ...domain.RoomFilterItem) ([]domain.Room, error) {
	var result *gorm.DB
	rooms := new([]types.Room)
	roomDomain := new([]domain.Room)
	offset := (pageIndex - 1) * pageSize
	if len(filter) > 0 {
		result = r.db.WithContext(ctx).Limit(int(pageSize)).Offset(int(offset)).Where(&filter[0]).Find(rooms)
	} else {
		result = r.db.WithContext(ctx).Limit(int(pageSize)).Offset(int(offset)).Where(&filter[0]).Find(rooms)

	}
	copier.Copy(roomDomain, rooms)
	return *roomDomain, result.Error
}
func (r *roomRepo) Update(ctx context.Context, UUID domain.RoomID, newRecord domain.Room) error {
	room := new(types.Room)
	copier.Copy(room, &newRecord)
	room.UpdateAt = time.Now()
	result := r.db.WithContext(ctx).Model(&room).Where("id = ?", UUID.String()).Updates(room)
	return result.Error
}
func (r *roomRepo) Delete(ctx context.Context, UUID domain.RoomID) error {
	room := new(types.Room)
	room.DeletedAt = time.Now()
	result := r.db.WithContext(ctx).Model(room).Where("id=?", UUID.String()).Updates(room)
	return result.Error
}
