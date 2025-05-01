package gateway

import (
	"bank_service/internal/common/types"
	"context"
)

type PaymentGateway interface {
	InitiatePayment(ctx context.Context, amount *types.Money, metadata map[string]interface{}) (string, error)
	VerifyPayment(ctx context.Context, referenceID string, amount int64) (bool, error)
	RefundPayment(ctx context.Context, referenceID string, amount *types.Money) error
	GetPaymentStatus(ctx context.Context, referenceID string) (string, error)
}
