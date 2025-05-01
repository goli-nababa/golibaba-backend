package rate

import (
	"context"
	"hotels-service/internal/rate/domain"
	"hotels-service/internal/rate/port"
	"time"
)

type rateService struct {
	repo port.Repo
}

func NewService(repo port.Repo) port.Service {
	return &rateService{
		repo: repo,
	}
}

func (s *rateService) CreateNewRate(ctx context.Context, rate domain.Rate) (domain.RateID, error) {
	rate.CreateAt = time.Now()
	return s.repo.Create(ctx, rate)
}

func (s *rateService) GetRateByID(ctx context.Context, UUID domain.RateID) (*domain.Rate, error) {
	return s.repo.GetByID(ctx, UUID)
}

func (s *rateService) GetAllRate(ctx context.Context, pageIndex, pageSize uint) ([]domain.Rate, error) {
	return s.repo.Get(ctx, pageIndex, pageSize)
}

func (s *rateService) FindRate(ctx context.Context, filters domain.RateFilterItem, pageIndex, pageSize uint) ([]domain.Rate, error) {
	return s.repo.Get(ctx, pageIndex, pageSize, filters)
}

func (s *rateService) EditeRate(ctx context.Context, UUID domain.RateID, newRate domain.Rate) error {
	existingRate, err := s.repo.GetByID(ctx, UUID)
	if err != nil {
		return err
	}

	newRate.CreateAt = existingRate.CreateAt
	newRate.ID = UUID

	return s.repo.Update(ctx, UUID, newRate)
}

func (s *rateService) DeleteRate(ctx context.Context, UUID domain.RateID) error {
	return s.repo.Delete(ctx, UUID)
}
