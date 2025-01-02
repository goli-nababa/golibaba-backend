package mapper

import (
	storageTypes "user_service/pkg/adapters/storage/types"

	"github.com/goli-nababa/golibaba-backend/common"
	"github.com/google/uuid"
)

func UserToStorage(user *common.User) *storageTypes.User {
	if user == nil {
		return nil
	}

	return &storageTypes.User{
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

func UserFromStorage(user *storageTypes.User) *common.User {
	if user == nil {
		return nil
	}

	uuid, _ := uuid.Parse(user.UUID)
	return &common.User{
		ID:        common.UserID(user.ID),
		UUID:      uuid,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  user.Password,
		Phone:     user.Phone,
		Blocked:   user.Blocked,
		WalletID:  common.WalletID(user.WalletID),
	}
}
