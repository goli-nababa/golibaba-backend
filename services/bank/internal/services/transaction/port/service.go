package port

import (
	moneyDomain "bank_service/internal/common/types"
	"bank_service/internal/services/transaction/domain"
	walletDomain "bank_service/internal/services/wallet/domain"
	"context"
	"time"
)

type TransactionService interface {
	CreateTransaction(ctx context.Context, req *TransactionRequest) (*domain.Transaction, error)
	ProcessTransaction(ctx context.Context, id domain.TransactionID) error
	CancelTransaction(ctx context.Context, id domain.TransactionID, reason string) error
	GetTransaction(ctx context.Context, id domain.TransactionID) (*domain.Transaction, error)
	ListTransactions(ctx context.Context, filter *TransactionFilter) ([]*domain.Transaction, error)
}

type TransactionRequest struct {
	FromWalletID walletDomain.WalletID
	ToWalletID   walletDomain.WalletID
	Amount       *moneyDomain.Money
	Type         domain.TransactionType
	Description  string
}

type TransactionFilter struct {
	WalletID walletDomain.WalletID
	Status   []domain.TransactionStatus
	Types    []domain.TransactionType
	From     time.Time
	To       time.Time
}
