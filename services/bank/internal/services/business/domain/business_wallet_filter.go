package domain

import (
	moneyDomain "bank_service/internal/common/types"
	walletDomain "bank_service/internal/services/wallet/domain"
	"time"
)

type BusinessWalletFilter struct {
	BusinessType   *BusinessType
	CommissionRate *float64
	Status         *walletDomain.WalletStatus
	MinBalance     *moneyDomain.Money
	MaxBalance     *moneyDomain.Money
	CreatedAfter   *time.Time
	CreatedBefore  *time.Time
	Limit          int
	Offset         int
}
