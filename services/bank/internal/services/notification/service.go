package notification

import (
	moneyDomain "bank_service/internal/common/types"
	"bank_service/internal/services/notification/port"
	"bank_service/internal/services/transaction/domain"
	walletDomain "bank_service/internal/services/wallet/domain"

	"context"
)

type Service struct {
}

func (s *Service) NotifyPaymentFailure(ctx context.Context, txID domain.TransactionID, userID uint64, reason string) error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) NotifyRefundProcessed(ctx context.Context, refundID string, userID uint64) error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) NotifyLowBalance(ctx context.Context, walletID walletDomain.WalletID, balance *moneyDomain.Money) error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) NotifyAdmins(ctx context.Context, subject string, data interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (s *Service) NotifyWalletCredit(ctx context.Context, userID uint64, amount *moneyDomain.Money) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) NotifyWalletDebit(ctx context.Context, userID uint64, amount *moneyDomain.Money) {
	//TODO implement me
	panic("implement me")
}

func NewService() port.Service {
	return &Service{}
}

func (s *Service) NotifyPaymentSuccess(ctx context.Context, txID domain.TransactionID, userID uint64) error {
	//TODO implement me
	panic("implement me")
}
