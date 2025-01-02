package out

import (
	moneyDomain "bank_service/internal/common/types"
	transactionDomain "bank_service/internal/services/transaction/domain"
	"context"
	"time"
)

type UserWalletService interface {
	ReserveMoney(ctx context.Context, userID string, amount *moneyDomain.Money) error
	ReleaseMoney(ctx context.Context, userID string, amount *moneyDomain.Money) error
	GetUserBalance(ctx context.Context, userID string) (*moneyDomain.Money, error)
}

type HotelPaymentService interface {
	ProcessRoomPayment(ctx context.Context, hotelID string, roomID string, userID string, amount *moneyDomain.Money, duration time.Duration) error
	ProcessRefund(ctx context.Context, bookingID string, refundReason string) error
	GetHotelTransactions(ctx context.Context, hotelID string, from, to time.Time) ([]*transactionDomain.Transaction, error)
}

type CommissionService interface {
	CalculateCommission(ctx context.Context, amount *moneyDomain.Money, businessType string) (*moneyDomain.Money, error)
	ProcessCommissionPayment(ctx context.Context, businessID string, transactionID string) error
	DistributeCommission(ctx context.Context, commissionID string) error
}
