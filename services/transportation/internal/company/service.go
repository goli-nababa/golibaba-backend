package company

import (
	"context"
	"log"
	"transportation/internal/company/domain"
	"transportation/internal/company/port"
)

type service struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &service{
		repo: repo,
	}
}
func (s *service) CreateCompany(ctx context.Context, dCompany domain.Company) (*domain.Company, error) {

	company, err := s.repo.Create(ctx, dCompany)
	if err != nil {
		log.Println("error on creating new company : ", err.Error())
		return nil, err
	}

	return company, nil
}
