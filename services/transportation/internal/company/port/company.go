package port

import (
	"context"
	commonDomain "transportation/internal/common/domain"
	"transportation/internal/company/domain"
)

type Repo interface {
	Create(ctx context.Context, company domain.Company) (*domain.Company, error)
	Update(ctx context.Context, id domain.CompanyId, companyDomain domain.Company) (*domain.Company, error)
	GetByID(ctx context.Context, id domain.CompanyId) (*domain.Company, error)
	Delete(ctx context.Context, id domain.CompanyId) error
	Get(ctx context.Context, request *commonDomain.RepositoryRequest) ([]domain.Company, error)

	CreateTechnicalTeam(ctx context.Context, techDomain domain.TechnicalTeam) (*domain.TechnicalTeam, error)
	UpdateTechnicalTeam(ctx context.Context, id domain.TechnicalTeamId, techDomain domain.TechnicalTeam) (*domain.TechnicalTeam, error)
	GetTechnicalTeamByID(ctx context.Context, id domain.TechnicalTeamId) (*domain.TechnicalTeam, error)
	DeleteTechnicalTeam(ctx context.Context, id domain.TechnicalTeamId) error
	GetTechnicalTeams(ctx context.Context, request *commonDomain.RepositoryRequest) ([]domain.TechnicalTeam, error)

	CreateTechnicalTeamMember(ctx context.Context, memberD domain.TechnicalTeamMember) (*domain.TechnicalTeamMember, error)
	DeleteTechnicalTeamMember(ctx context.Context, id domain.TechnicalTeamMemberId) error
	GetTechnicalTeamMembers(ctx context.Context, request *commonDomain.RepositoryRequest) ([]domain.TechnicalTeamMember, error)
}
