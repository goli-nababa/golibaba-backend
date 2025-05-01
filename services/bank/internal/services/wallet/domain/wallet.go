package domain

import (
	"bank_service/internal/common/types"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInsufficientFunds      = errors.New("insufficient funds")
	ErrInvalidStatus          = errors.New("invalid wallet status")
	ErrConcurrentModification = errors.New("wallet modified concurrently")
	ErrInsufficientBalance    = errors.New("insufficient balance")
	ErrInvalidWalletStatus    = errors.New("invalid wallet status")
	ErrWalletLocked           = errors.New("wallet is locked")
	ErrWalletBlocked          = errors.New("wallet is blocked")
	ErrBelowMinimumBalance    = errors.New("balance would fall below minimum")
	ErrInvalidAmount          = errors.New("invalid amount")
	ErrDailyLimitExceeded     = errors.New("daily transaction limit exceeded")
	ErrMonthlyLimitExceeded   = errors.New("monthly transaction limit exceeded")
	ErrConcurrentOperation    = errors.New("concurrent operation exceeded")
)

type WalletID string

type WalletStatus string

const (
	WalletStatusActive   WalletStatus = "active"
	WalletStatusInactive WalletStatus = "inactive"
	WalletStatusBlocked  WalletStatus = "blocked"
	WalletStatusLocked   WalletStatus = "locked"
)

type WalletType string

const (
	WalletTypePersonal        WalletType = "personal"
	WalletTypeBusinessBasic   WalletType = "business_basic"
	WalletTypeBusinessPremium WalletType = "business_premium"
)

type Wallet struct {
	ID                  WalletID
	UserID              uint64
	Balance             *types.Money
	Status              WalletStatus
	CreatedAt           time.Time
	UpdatedAt           time.Time
	Version             int
	Type                WalletType
	Locks               []WalletLock
	DailyTransactions   map[time.Time]*types.Money // Key is date
	MonthlyTransactions map[time.Time]*types.Money // Key is month start
}

type WalletLock struct {
	Reason    string
	ExpiresAt time.Time
	CreatedAt time.Time
}

func NewWallet(userID uint64, walletType WalletType, currency string) (*Wallet, error) {
	initialBalance, err := types.NewMoney(0, currency)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	return &Wallet{
		ID:      WalletID(uuid.New().String()),
		UserID:  userID,
		Type:    walletType,
		Status:  WalletStatusActive,
		Balance: initialBalance,

		DailyTransactions:   make(map[time.Time]*types.Money),
		MonthlyTransactions: make(map[time.Time]*types.Money),
		Locks:               make([]WalletLock, 0),
		CreatedAt:           now,
		UpdatedAt:           now,
		Version:             1,
	}, nil
}

func (w *Wallet) AdjustBalance(amount *types.Money) error {
	if amount.Currency != w.Balance.Currency {
		return types.ErrCurrencyMismatch
	}
	if w.Status != WalletStatusActive {
		return ErrInvalidStatus
	}
	newBalance, err := w.Balance.Add(amount)
	if err != nil {
		return err
	}
	if newBalance.Amount < 0 {
		return ErrInsufficientFunds
	}
	w.Balance = newBalance
	w.UpdatedAt = time.Now()
	return nil
}

func (w *Wallet) Credit(amount *types.Money) error {
	if err := w.validateStatus(); err != nil {
		return err
	}

	if err := w.validateAmount(amount); err != nil {
		return err
	}

	w.Balance, _ = w.Balance.Add(amount)
	w.updateTransactionLimits(amount)
	w.UpdatedAt = time.Now()
	return nil
}

func (w *Wallet) Debit(amount *types.Money) error {
	if err := w.validateStatus(); err != nil {
		return err
	}

	if err := w.validateAmount(amount); err != nil {
		return err
	}

	remaining, err := w.Balance.Subtract(amount)
	if err != nil {
		return err
	}

	w.Balance = remaining
	w.updateTransactionLimits(amount)
	w.UpdatedAt = time.Now()
	return nil
}

func (w *Wallet) Block(reason string) error {
	if w.Status == WalletStatusBlocked {
		return nil
	}

	w.Status = WalletStatusBlocked
	w.UpdatedAt = time.Now()

	return nil
}

func (w *Wallet) Lock(reason string, duration time.Duration) error {
	if w.Status == WalletStatusBlocked {
		return ErrWalletBlocked
	}

	now := time.Now()
	w.Locks = append(w.Locks, WalletLock{
		Reason:    reason,
		ExpiresAt: now.Add(duration),
		CreatedAt: now,
	})

	w.Status = WalletStatusLocked
	w.UpdatedAt = now

	return nil
}

func (w *Wallet) Unlock() error {
	if w.Status == WalletStatusBlocked {
		return ErrWalletBlocked
	}

	w.Locks = nil
	w.Status = WalletStatusActive
	w.UpdatedAt = time.Now()

	return nil
}

func (w *Wallet) validateStatus() error {
	switch w.Status {
	case WalletStatusActive:
		return nil
	case WalletStatusBlocked:
		return ErrWalletBlocked
	case WalletStatusLocked:
		return ErrWalletLocked
	default:
		return ErrInvalidWalletStatus
	}
}

func (w *Wallet) validateAmount(amount *types.Money) error {
	if amount == nil || amount.Amount <= 0 {
		return ErrInvalidAmount
	}
	if amount.Currency != w.Balance.Currency {
		return types.ErrCurrencyMismatch
	}
	return nil
}

func (w *Wallet) updateTransactionLimits(amount *types.Money) {
	today := time.Now().Truncate(24 * time.Hour)
	monthStart := time.Date(today.Year(), today.Month(), 1, 0, 0, 0, 0, time.UTC)

	if daily, exists := w.DailyTransactions[today]; exists {
		w.DailyTransactions[today], _ = daily.Add(amount)
	} else {
		w.DailyTransactions[today] = amount
	}

	if monthly, exists := w.MonthlyTransactions[monthStart]; exists {
		w.MonthlyTransactions[monthStart], _ = monthly.Add(amount)
	} else {
		w.MonthlyTransactions[monthStart] = amount
	}

	w.cleanupTransactionHistory()
}

func (w *Wallet) cleanupTransactionHistory() {
	// Cleanup daily transactions older than 30 days
	thirtyDaysAgo := time.Now().AddDate(0, 0, -30).Truncate(24 * time.Hour)
	for date := range w.DailyTransactions {
		if date.Before(thirtyDaysAgo) {
			delete(w.DailyTransactions, date)
		}
	}

	// Cleanup monthly transactions older than 12 months
	twelveMonthsAgo := time.Now().AddDate(-1, 0, 0).Truncate(24 * time.Hour)
	for date := range w.MonthlyTransactions {
		if date.Before(twelveMonthsAgo) {
			delete(w.MonthlyTransactions, date)
		}
	}
}
