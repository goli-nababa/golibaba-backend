package helpers

import (
	"github.com/goli-nababa/golibaba-backend/common"
	"github.com/google/uuid"
	"time"
	"user_service/api/http/helpers"
	"user_service/api/http/types"
)

func RegisterRequestToUserDomain(req types.RegisterRequest) (*domain.User, error) {
	birthday, err := helpers.IsValidDate(req.Birthday)

	if err != nil {
		return nil, errors.New("birthday invalid")
	}

	return &common.User{
		ID:        0,
		UUID:      uuid.UUID{},
		FirstName: "",
		LastName:  "",
		Email:     "",
		Password:  "",
		Phone:     "",
		Blocked:   false,
		WalletID:  0,
		Role:      "",
	}, nil
}
