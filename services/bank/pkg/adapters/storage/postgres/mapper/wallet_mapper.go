package mapper

import (
	businessWallet "bank_service/internal/services/business/domain"
	"bank_service/internal/services/wallet/domain"
	"bank_service/pkg/adapters/storage/postgres/model"
	"encoding/json"
)

func WalletDomainToModel(w *domain.Wallet) *model.WalletModel {
	return &model.WalletModel{
		ID:        string(w.ID),
		UserID:    w.UserID,
		Balance:   w.Balance.Amount,
		Currency:  w.Balance.Currency,
		Status:    string(w.Status),
		CreatedAt: w.CreatedAt,
		UpdatedAt: w.UpdatedAt,
		Version:   w.Version,
	}
}

func BusinessWalletDomainToModel(w *businessWallet.BusinessWallet) (*model.BusinessWalletModel, error) {
	baseModel := WalletDomainToModel(w.Wallet)
	bankInfo, err := json.Marshal(w.BankInfo)
	if err != nil {
		return nil, err
	}

	return &model.BusinessWalletModel{
		WalletModel:       *baseModel,
		BusinessID:        w.BusinessID,
		BusinessType:      string(w.BusinessType),
		CommissionRate:    w.CommissionRate,
		PayoutSchedule:    w.PayoutSchedule,
		LastPayoutDate:    w.LastPayoutDate,
		BankInfo:          bankInfo,
		MinPayoutAmount:   w.MinimumPayout.Amount,
		MinPayoutCurrency: w.MinimumPayout.Currency,
	}, nil
}
