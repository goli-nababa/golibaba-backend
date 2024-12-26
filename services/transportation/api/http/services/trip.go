package services

import (
	"context"
	"errors"
	companyDomain "transportation/internal/company/domain"
	companyPort "transportation/internal/company/port"

	"transportation/internal/trip/domain"

	"transportation/internal/trip/port"
)

type TripService struct {
	MainService    port.Service
	companyService companyPort.Service
}

func NewTripService(svc port.Service, companySvc companyPort.Service) *TripService {
	return &TripService{
		MainService:    svc,
		companyService: companySvc,
	}
}

func (s *TripService) CreateTrip(ctx context.Context, req domain.CreateTripRequest) (domain.Trip, error) {
	userId := 1
	companies, err := s.companyService.GetCompanies(ctx, companyDomain.CompanyFilter{OwnerId: uint(userId)})
	if err != nil {
		return domain.Trip{}, err
	}

	ownCompany := false
	for _, v := range companies {
		if v.ID == req.CompanyId {
			ownCompany = true
		}
	}

	if !ownCompany {
		return domain.Trip{}, errors.New("this company is not for this user")
	}

	return s.MainService.CreateTrip(ctx, req)

}

func (s *TripService) ConfirmTechnicalTeam(ctx context.Context, id domain.TripId) (domain.Trip, error) {

	userId := 1
	trip, err := s.MainService.GetTrip(ctx, id)

	if err != nil {
		return domain.Trip{}, err
	}
	teamMembers, err := s.companyService.GetTechnicalTeamMembers(ctx, companyDomain.TechnicalTeamMemberFilter{TechnicalTeamId: companyDomain.TechnicalTeamId(trip.TechTeamId), MemberId: companyDomain.MemberId(userId)})

	if err != nil {
		return trip, err
	}
	if len(teamMembers) < 1 || teamMembers == nil {

		return trip, errors.New("user is not member of trip technical team")
	}

	return s.MainService.ConfirmTechnicalTeam(ctx, id)
}

func (s *TripService) EndTrip(ctx context.Context, id domain.TripId) (domain.Trip, error) {

	userId := 1
	trip, err := s.MainService.GetTrip(ctx, id)

	if err != nil {
		return domain.Trip{}, err
	}
	teamMembers, err := s.companyService.GetTechnicalTeamMembers(ctx, companyDomain.TechnicalTeamMemberFilter{TechnicalTeamId: companyDomain.TechnicalTeamId(trip.TechTeamId), MemberId: companyDomain.MemberId(userId)})
	if err != nil {
		return trip, err
	}
	if len(teamMembers) < 1 || teamMembers == nil || userId != int(trip.Company.OwnerId) {
		return trip, errors.New("user is not member of trip technical team")
	}
	return s.MainService.EndTrip(ctx, id)
}
