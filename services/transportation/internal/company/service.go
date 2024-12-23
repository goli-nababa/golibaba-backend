package company

import (
	"context"
	"strconv"
	commonDomain "transportation/internal/common/domain"
	"transportation/internal/company/domain"

	"transportation/internal/company/port"
	"transportation/pkg/logging"
)

type service struct {
	repo   port.Repo
	logger logging.Logger
}

func NewService(repo port.Repo, logger logging.Logger) port.Service {
	return &service{
		repo:   repo,
		logger: logger,
	}
}
func (s *service) CreateCompany(ctx context.Context, dCompany domain.Company) (*domain.Company, error) {

	company, err := s.repo.Create(ctx, dCompany)
	if err != nil {
		s.logger.Error(logging.Internal, logging.FailedToCreateCompany, "error in create company", map[logging.ExtraKey]interface{}{logging.ErrorMessage: err.Error()})
		return nil, err
	}

	company, err = s.repo.GetByID(ctx, company.ID, "TransportationType")
	if err != nil {
		s.logger.Error(logging.Internal, logging.FailedToCreateCompany, "error in create company", map[logging.ExtraKey]interface{}{logging.ErrorMessage: err.Error()})
		return nil, err
	}
	return company, err
}

func (s *service) UpdateCompany(ctx context.Context, companyId domain.CompanyId, company domain.Company) (*domain.Company, error) {

	response, err := s.repo.Update(ctx, companyId, company)

	if err != nil {
		s.logger.Error(logging.Internal, logging.FailedToUpdateCompany, "error in update company", map[logging.ExtraKey]interface{}{logging.ErrorMessage: err.Error()})
		return nil, err
	}
	return response, nil
}
func (s *service) DeleteCompany(ctx context.Context, companyId domain.CompanyId) error {
	err := s.repo.Delete(ctx, companyId)
	if err != nil {
		s.logger.Error(logging.Internal, logging.FailedToDeleteCompany, "error in delete company", map[logging.ExtraKey]interface{}{logging.ErrorMessage: err.Error()})
		return err
	}
	return nil
}

func (s *service) GetCompanies(ctx context.Context, filters domain.CompanyFilter) ([]domain.Company, error) {
	rFilters := []*commonDomain.RepositoryFilter{
		&commonDomain.RepositoryFilter{Field: "owner_id", Operator: "=", Value: strconv.Itoa(int(filters.OwnerId))}}

	if filters.TransportationTypeId > 0 {
		rFilters = append(rFilters,
			&commonDomain.RepositoryFilter{Field: "transportation_type_id", Operator: "=", Value: strconv.Itoa(int(filters.TransportationTypeId))})
	}
	companies, err := s.repo.Get(ctx,
		&commonDomain.RepositoryRequest{Filters: rFilters, Preloads: []string{"TransportationType", "TechnicalTeams", "TechnicalTeams.TechnicalTeamMembers"}})
	if err != nil {
		s.logger.Error(logging.Internal, logging.FailedToGetCompanies, "error in GetCompanies company", map[logging.ExtraKey]interface{}{logging.ErrorMessage: err.Error()})
		return []domain.Company{}, err
	}

	return companies, nil
}

func (s *service) CreateTechnicalTeam(ctx context.Context, team domain.TechnicalTeam) (*domain.TechnicalTeam, error) {
	return s.repo.CreateTechnicalTeam(ctx, team)
}
func (s *service) UpdateTechnicalTeam(ctx context.Context, teamId domain.TechnicalTeamId, team domain.TechnicalTeam) (*domain.TechnicalTeam, error) {
	return s.repo.UpdateTechnicalTeam(ctx, teamId, team)
}
func (s *service) DeleteTechnicalTeam(ctx context.Context, teamId domain.TechnicalTeamId) error {
	return s.repo.DeleteTechnicalTeam(ctx, teamId)
}
func (s *service) GetTechnicalTeams(ctx context.Context, companyId domain.CompanyId) ([]domain.TechnicalTeam, error) {
	return s.repo.GetTechnicalTeams(ctx, &commonDomain.RepositoryRequest{Filters: []*commonDomain.RepositoryFilter{&commonDomain.RepositoryFilter{Field: "company_id", Operator: "=", Value: strconv.Itoa(int(companyId))}}})

}

func (s *service) AddMembersToTechnicalTeam(ctx context.Context, teamId domain.TechnicalTeamId, memberIds []domain.MemberId) (*domain.TechnicalTeam, error) {

	team, err := s.repo.GetTechnicalTeamByID(ctx, teamId)
	if err != nil {
		return nil, err
	}
	for _, v := range memberIds {

		_, err := s.repo.CreateTechnicalTeamMember(ctx, domain.TechnicalTeamMember{MemberId: v, TechnicalTeamId: teamId})
		if err != nil {
			return team, err
		}
	}
	return team, nil
}
func (s *service) RemoveMembersFromTechnicalTeam(ctx context.Context, teamId domain.TechnicalTeamId, technicalTeamMemberId domain.TechnicalTeamMemberId) error {
	return s.repo.DeleteTechnicalTeamMember(ctx, technicalTeamMemberId)
}
