package port

import (
	moneyDomain "bank_service/internal/common/types"
	domain2 "bank_service/internal/services/payment/domain"
	"bank_service/internal/services/transaction/domain"
	walletDomain "bank_service/internal/services/wallet/domain"
	"context"
)

type Service interface {
	ProcessPayment(ctx context.Context, tx *domain.Transaction) (string, error)
	ProcessRefund(ctx context.Context, tx *domain.Transaction) error
	VerifyPayment(ctx context.Context, referenceID string) (bool, error)

	RefundPayment(ctx context.Context, paymentID string) error

	ProcessCommission(ctx context.Context, amount *moneyDomain.Money, businessID uint64) error

	ChargeWallet(ctx context.Context, walletID walletDomain.WalletID, amount *moneyDomain.Money) error
	WithdrawFromWallet(ctx context.Context, walletID walletDomain.WalletID, amount *moneyDomain.Money) error

	GetTransactionHistory(ctx context.Context, walletID walletDomain.WalletID) ([]*domain.Transaction, error)

	HandlePaymentCallback(ctx context.Context, callback *domain2.PaymentCallback) (bool, string, error)
}
