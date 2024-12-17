package services

import companyService "transportation/internal/company/port"

type CompanyService struct {
	svc companyService.Service
}

func NewCompanyService(svc companyService.Service) *CompanyService {
	return &CompanyService{
		svc: svc,
	}
}
