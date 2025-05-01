package service

import (
	businessDomain "bank_service/internal/services/business/domain"
	"bank_service/internal/services/commission/port"
	"context"
)

type rateProvider struct {
}

func NewCommissionRateProvider() port.CommissionRateProvider {
	return &rateProvider{}
}

func (p *rateProvider) GetRate(ctx context.Context, businessType businessDomain.BusinessType) (float64, error) {
	switch businessType {
	case businessDomain.BusinessTypeHotel:
		return 0.10, nil // 10%
	case businessDomain.BusinessTypeAirline:
		return 0.08, nil // 8%
	case businessDomain.BusinessTypeTravelAgency:
		return 0.05, nil // 5%
	case businessDomain.BusinessTypeShip:
		return 0.07, nil // 7%
	case businessDomain.BusinessTypeTrain:
		return 0.06, nil // 6%
	case businessDomain.BusinessTypeBus:
		return 0.04, nil // 4%
	default:
		return 0.10, nil // Default 10%
	}
}
