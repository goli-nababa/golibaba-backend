package port

import (
	"context"
	"transportation/internal/company/domain"
)

type Service interface {
	CreateCompany(ctx context.Context, company domain.Company) (*domain.Company, error)
	UpdateCompany(ctx context.Context, companyId domain.CompanyId, company domain.Company) (*domain.Company, error)
	DeleteCompany(ctx context.Context, companyId domain.CompanyId) error
	GetCompanies(ctx context.Context, company domain.CompanyFilter) ([]domain.Company, error)

	CreateTechnicalTeam(ctx context.Context, team domain.TechnicalTeam) (*domain.TechnicalTeam, error)
	UpdateTechnicalTeam(ctx context.Context, teamId domain.TechnicalTeamId, team domain.TechnicalTeam) (*domain.TechnicalTeam, error)
	DeleteTechnicalTeam(ctx context.Context, teamId domain.TechnicalTeamId) error
	GetTechnicalTeams(ctx context.Context, companyId domain.CompanyId) ([]domain.TechnicalTeam, error)

	AddMembersToTechnicalTeam(ctx context.Context, teamId domain.TechnicalTeamId, memberIds []domain.MemberId) (*domain.TechnicalTeam, error)
	RemoveMembersFromTechnicalTeam(ctx context.Context, teamId domain.TechnicalTeamId, technicalTeamMemberId domain.TechnicalTeamMemberId) error
	GetTechnicalTeamMembers(ctx context.Context, filters domain.TechnicalTeamMemberFilter) ([]domain.TechnicalTeamMember, error)
}
