package mapper

import (
	"github.com/goli-nababa/golibaba-backend/common"
	"github.com/google/uuid"
	"user_service/pkg/adapters/storage/types"
)

func UserToStorage(user *common.User) *types.User {
	if user == nil {
		return nil
	}

	return &types.User{
		ID:        uint(user.ID),
		UUID:      user.UUID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  user.Password,
		Phone:     user.Phone,
		Blocked:   user.Blocked,
		WalletID:  uint(user.WalletID),
	}
}

func UserFromStorage(user *types.User) *common.User {
	if user == nil {
		return nil
	}

	parsedUUID, _ := uuid.Parse(user.UUID)
	return &common.User{
		ID:        common.UserID(user.ID),
		UUID:      parsedUUID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  user.Password,
		Phone:     user.Phone,
		Blocked:   user.Blocked,
		WalletID:  common.WalletID(user.WalletID),
	}
}
