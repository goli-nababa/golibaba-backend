package storage

import (
	"context"
	"fmt"

	"github.com/goli-nababa/golibaba-backend/common"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *common.User) error {
	return nil
}
