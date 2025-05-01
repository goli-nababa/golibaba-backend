package model

import (
	moneyDomain "bank_service/internal/common/types"
	"bank_service/internal/services/wallet/domain"
	"time"
)

type WalletModel struct {
	ID        string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID    uint64
	Balance   int64
	Currency  string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
	Version   int
	tableName struct{} `gorm:"table:wallets"`
}

func (m *WalletModel) ToDomain() *domain.Wallet {
	return &domain.Wallet{
		ID:        domain.WalletID(m.ID),
		UserID:    m.UserID,
		Balance:   &moneyDomain.Money{Amount: m.Balance, Currency: m.Currency},
		Status:    domain.WalletStatus(m.Status),
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		Version:   m.Version,
	}
}
