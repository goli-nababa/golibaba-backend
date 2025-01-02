package services

import (
	"context"
	"transportation/api/http/types"
	"transportation/internal/company/domain"
	companyService "transportation/internal/company/port"
)

type CompanyService struct {
	svc companyService.Service
}

func NewCompanyService(svc companyService.Service) *CompanyService {
	return &CompanyService{
		svc: svc,
	}
}

func (s *CompanyService) CreateCompany(ctx context.Context, req types.CreateCompanyRequest) (types.CompanyResponse, error) {
	d, err := s.svc.CreateCompany(ctx, domain.Company{Name: req.Name, OwnerId: 1, TransportationTypeId: req.TransportationTypeId})
	if err != nil {
		return types.CompanyResponse{}, err
	}

	return types.CompanyResponse{Id: d.ID, Name: d.Name, OwnerId: d.OwnerId, TransportationType: d.TransportationType}, nil
}

func (s *CompanyService) GetCompanies(ctx context.Context, req types.FilterCompaniesRequest) ([]domain.Company, error) {
	return s.svc.GetCompanies(ctx, domain.CompanyFilter{OwnerId: 1, TransportationTypeId: req.TransportationTypeId})
}

//....
