package port

import (
	"bank_service/internal/common/types"
	"bank_service/internal/services/analytics/domain"
	businessDomain "bank_service/internal/services/business/domain"
	txDomain "bank_service/internal/services/transaction/domain"
	walletDomain "bank_service/internal/services/wallet/domain"
	"context"
	"time"
)

type Service interface {
	CreateBusinessWallet(ctx context.Context, businessID uint64, businessType businessDomain.BusinessType, currency string) (*businessDomain.BusinessWallet, error)
	GetBusinessWallet(ctx context.Context, walletID walletDomain.WalletID) (*businessDomain.BusinessWallet, error)
	UpdateBusinessWallet(ctx context.Context, wallet *businessDomain.BusinessWallet) error

	ProcessCommission(ctx context.Context, tx *txDomain.Transaction) error
	SetCommissionRate(ctx context.Context, walletID walletDomain.WalletID, rate float64) error
	GetCommissionRate(ctx context.Context, walletID walletDomain.WalletID) (float64, error)

	RequestPayout(ctx context.Context, walletID walletDomain.WalletID, amount *types.Money) error
	SetPayoutSchedule(ctx context.Context, walletID walletDomain.WalletID, schedule string) error
	SetMinimumPayout(ctx context.Context, walletID walletDomain.WalletID, amount *types.Money) error

	UpdateBankInfo(ctx context.Context, walletID walletDomain.WalletID, bankInfo *businessDomain.BankAccountInfo) error
	GetBankInfo(ctx context.Context, walletID walletDomain.WalletID) (*businessDomain.BankAccountInfo, error)

	GetBusinessStats(ctx context.Context, businessID uint64, startDate, endDate time.Time) (*domain.BusinessStats, error)
	GetCommissionHistory(ctx context.Context, businessID uint64, startDate, endDate time.Time) ([]*domain.CommissionEntry, error)
	GetPayoutHistory(ctx context.Context, walletID walletDomain.WalletID, startDate, endDate time.Time) ([]*businessDomain.PayoutEntry, error)
}
