package domain

import (
	"bank_service/internal/common/types"
)

type WalletEvent interface {
	EventType() string
}

type WalletCreatedEvent struct {
	Wallet *Wallet
}

func (e WalletCreatedEvent) EventType() string {
	return "wallet.created"
}

type WalletStatusChangedEvent struct {
	WalletID  WalletID
	OldStatus WalletStatus
	NewStatus WalletStatus
	Reason    string
}

func (e WalletStatusChangedEvent) EventType() string {
	return "wallet.status.changed"
}

type WalletBalanceChangedEvent struct {
	WalletID     WalletID
	OldBalance   *types.Money
	NewBalance   *types.Money
	ChangeAmount *types.Money
	ChangeType   string // "credit" or "debit"
}

func (e WalletBalanceChangedEvent) EventType() string {
	return "wallet.balance.changed"
}

type WalletLimitExceededEvent struct {
	WalletID  WalletID
	LimitType string // "daily" or "monthly"
	Attempted *types.Money
	Limit     *types.Money
}

func (e WalletLimitExceededEvent) EventType() string {
	return "wallet.limit.exceeded"
}

type LowBalanceEvent struct {
	WalletID       WalletID
	Balance        *types.Money
	MinimumBalance *types.Money
}

func (e LowBalanceEvent) EventType() string {
	return "wallet.balance.low"
}
