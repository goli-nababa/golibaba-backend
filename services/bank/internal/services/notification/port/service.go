package port

import (
	moneyDomain "bank_service/internal/common/types"
	transactionDomain "bank_service/internal/services/transaction/domain"
	walletDomain "bank_service/internal/services/wallet/domain"

	"context"
)

type Service interface {
	NotifyPaymentSuccess(ctx context.Context, txID transactionDomain.TransactionID, userID uint64) error
	NotifyPaymentFailure(ctx context.Context, txID transactionDomain.TransactionID, userID uint64, reason string) error
	NotifyRefundProcessed(ctx context.Context, refundID string, userID uint64) error
	NotifyLowBalance(ctx context.Context, walletID walletDomain.WalletID, balance *moneyDomain.Money) error

	NotifyAdmins(ctx context.Context, subject string, data interface{}) error

	NotifyWalletCredit(ctx context.Context, userID uint64, amount *moneyDomain.Money)
	NotifyWalletDebit(ctx context.Context, userID uint64, amount *moneyDomain.Money)
}
