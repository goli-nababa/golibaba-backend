package storage

import (
	"context"
	commonDomain "transportation/internal/common/domain"
	"transportation/internal/company/domain"
	"transportation/internal/company/port"
	"transportation/pkg/adapters/storage/mapper"
	"transportation/pkg/adapters/storage/types"

	"gorm.io/gorm"
)

type companyRepo struct {
	db *gorm.DB
}

func NewCompanyRepo(db *gorm.DB) port.Repo {
	repo := &companyRepo{db}
	return repo
}

func (r *companyRepo) Create(ctx context.Context, companyDomain domain.Company) (*domain.Company, error) {
	company := types.Company{}
	if err := mapper.ConvertTypes(companyDomain, &company); err != nil {
		return nil, err
	}

	if err := CreateRecord(r.db, &company); err != nil {
		return nil, err
	}

	if err := mapper.ConvertTypes(company, &companyDomain); err != nil {
		return nil, err
	}
	return &companyDomain, nil
}

func (r *companyRepo) Update(ctx context.Context, id domain.CompanyId, companyDomain domain.Company) (*domain.Company, error) {
	company := types.Company{}
	if err := mapper.ConvertTypes(companyDomain, &company); err != nil {
		return nil, err
	}
	if err := UpdateRecord(r.db, id, company); err != nil {
		return nil, err
	}

	domain, err := r.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return domain, nil
}

func (r *companyRepo) GetByID(ctx context.Context, id domain.CompanyId) (*domain.Company, error) {
	company, err := GetRecordByID[types.Company](r.db, id, nil)
	if err != nil {
		return nil, err
	}
	if company == nil {
		return nil, gorm.ErrRecordNotFound
	}

	companyDomain := domain.Company{}
	if err := mapper.ConvertTypes(*company, &companyDomain); err != nil {
		return nil, err
	}

	return &companyDomain, nil
}

func (r *companyRepo) Delete(ctx context.Context, id domain.CompanyId) error {
	return DeleteRecordByID[types.Company](r.db, id)
}

func (r *companyRepo) Get(ctx context.Context, request *commonDomain.RepositoryRequest) ([]domain.Company, error) {
	companies, err := GetRecords[types.Company](r.db, request)
	if err != nil {
		return []domain.Company{}, err
	}

	companyDomains := []domain.Company{}

	mapper.ConvertTypes(companies, &companyDomains)
	return companyDomains, nil
}

func (r *companyRepo) CreateTechnicalTeam(ctx context.Context, techDomain domain.TechnicalTeam) (*domain.TechnicalTeam, error) {
	techTeam := types.TechnicalTeam{}
	if err := mapper.ConvertTypes(techDomain, &techTeam); err != nil {
		return nil, err
	}

	if err := CreateRecord(r.db, &techTeam); err != nil {
		return nil, err
	}

	if err := mapper.ConvertTypes(techTeam, &techDomain); err != nil {
		return nil, err
	}
	return &techDomain, nil
}

func (r *companyRepo) UpdateTechnicalTeam(ctx context.Context, id domain.TechnicalTeamId, techDomain domain.TechnicalTeam) (*domain.TechnicalTeam, error) {
	technicalTeam := types.TechnicalTeam{}
	if err := mapper.ConvertTypes(techDomain, &technicalTeam); err != nil {
		return nil, err
	}
	if err := UpdateRecord(r.db, id, technicalTeam); err != nil {
		return nil, err
	}

	domain, err := r.GetTechnicalTeamByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return domain, nil
}

func (r *companyRepo) GetTechnicalTeamByID(ctx context.Context, id domain.TechnicalTeamId) (*domain.TechnicalTeam, error) {
	team, err := GetRecordByID[types.TechnicalTeam](r.db, id, nil)
	if err != nil {
		return nil, err
	}
	if team == nil {
		return nil, gorm.ErrRecordNotFound
	}

	domain := domain.TechnicalTeam{}
	if err := mapper.ConvertTypes(*team, &domain); err != nil {
		return nil, err
	}

	return &domain, nil
}

func (r *companyRepo) DeleteTechnicalTeam(ctx context.Context, id domain.TechnicalTeamId) error {
	return DeleteRecordByID[types.TechnicalTeam](r.db, id)
}

func (r *companyRepo) GetTechnicalTeams(ctx context.Context, request *commonDomain.RepositoryRequest) ([]domain.TechnicalTeam, error) {
	companies, err := GetRecords[types.TechnicalTeam](r.db, request)
	if err != nil {
		return []domain.TechnicalTeam{}, err
	}

	teamsDomains := []domain.TechnicalTeam{}

	mapper.ConvertTypes(companies, &teamsDomains)
	return teamsDomains, nil
}

func (r *companyRepo) CreateTechnicalTeamMember(ctx context.Context, memberD domain.TechnicalTeamMember) (*domain.TechnicalTeamMember, error) {
	member := types.TechnicalTeamMember{}
	if err := mapper.ConvertTypes(memberD, &member); err != nil {
		return nil, err
	}

	if err := CreateRecord(r.db, &member); err != nil {
		return nil, err
	}

	if err := mapper.ConvertTypes(member, &memberD); err != nil {
		return nil, err
	}
	return &memberD, nil
}

func (r *companyRepo) DeleteTechnicalTeamMember(ctx context.Context, id domain.TechnicalTeamMemberId) error {
	return DeleteRecordByID[types.TechnicalTeamMember](r.db, id)
}

func (r *companyRepo) GetTechnicalTeamMembers(ctx context.Context, request *commonDomain.RepositoryRequest) ([]domain.TechnicalTeamMember, error) {
	members, err := GetRecords[types.TechnicalTeamMember](r.db, request)
	if err != nil {
		return []domain.TechnicalTeamMember{}, err
	}

	membersD := []domain.TechnicalTeamMember{}

	mapper.ConvertTypes(members, &membersD)
	return membersD, nil
}
