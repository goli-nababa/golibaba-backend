package domain

import (
	moneyDomain "bank_service/internal/common/types"
	walletDomain "bank_service/internal/services/wallet/domain"
	"errors"
	"time"
)

var (
	ErrInvalidBusinessType   = errors.New("invalid business type")
	ErrInvalidCommissionRate = errors.New("invalid commission rate")
)

const CENTRAL_WALLET_ID = "00000000-0000-0000-0000-000000000001"

type BusinessType string

const (
	BusinessTypeHotel        BusinessType = "hotel"
	BusinessTypeAirline      BusinessType = "airline"
	BusinessTypeTravelAgency BusinessType = "travel_agency"
	BusinessTypeShip         BusinessType = "ship"
	BusinessTypeTrain        BusinessType = "train"
	BusinessTypeBus          BusinessType = "bus"
)

type BankAccountInfo struct {
	AccountNumber string
	IBAN          string
	BankName      string
	AccountName   string
	CardNumber    string
}

type BusinessWallet struct {
	*walletDomain.Wallet // Embed base wallet
	BusinessID           uint64
	BusinessType         BusinessType
	CommissionRate       float64
	PayoutSchedule       string
	LastPayoutDate       *time.Time
	MinimumPayout        *moneyDomain.Money
	BankInfo             *BankAccountInfo
}

func NewBusinessWallet(businessID uint64, businessType BusinessType, currency string) (*BusinessWallet, error) {
	baseWallet, err := walletDomain.NewWallet(businessID, walletDomain.WalletTypeBusinessBasic, currency)
	if err != nil {
		return nil, err
	}

	return &BusinessWallet{
		Wallet:         baseWallet,
		BusinessID:     businessID,
		BusinessType:   businessType,
		CommissionRate: getDefaultCommissionRate(businessType),
	}, nil
}

func getDefaultCommissionRate(bt BusinessType) float64 {
	switch bt {
	case BusinessTypeHotel:
		return 0.1 // 10%
	case BusinessTypeAirline:
		return 0.08 // 8%
	case BusinessTypeTravelAgency:
		return 0.05 // 5%
	default:
		return 0.1 // 10% default
	}
}
