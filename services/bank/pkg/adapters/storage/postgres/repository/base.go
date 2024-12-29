package repository

import (
	"bank_service/pkg/transaction"
	"context"
	"gorm.io/gorm"
)

type BaseRepository struct {
	_db *gorm.DB
}

func NewBaseRepository(db *gorm.DB) BaseRepository {
	return BaseRepository{_db: db}
}

func (r *BaseRepository) DB(ctx context.Context) *gorm.DB {
	if tx := transaction.GetTransaction(ctx); tx != nil {
		return tx
	}
	return r._db
}
