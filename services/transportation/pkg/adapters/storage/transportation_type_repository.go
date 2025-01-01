package storage

import (
	"context"
	commonDomain "transportation/internal/common/domain"

	"transportation/internal/transportation_type/domain"
	"transportation/internal/transportation_type/port"
	"transportation/pkg/adapters/storage/mapper"
	"transportation/pkg/adapters/storage/types"

	"gorm.io/gorm"
)

type TransportationTypeRepo struct {
	db *gorm.DB
}

func NewTransportationTypeRepo(db *gorm.DB) port.Repo {
	repo := &TransportationTypeRepo{db}
	return repo
}

func (r *TransportationTypeRepo) Create(ctx context.Context, TransportationTypeDomain domain.TransportationType) (*domain.TransportationType, error) {
	TransportationType := types.TransportationType{}
	if err := mapper.ConvertTypes(TransportationTypeDomain, &TransportationType); err != nil {
		return nil, err
	}

	if err := CreateRecord(r.db, &TransportationType); err != nil {
		return nil, err
	}

	if err := mapper.ConvertTypes(TransportationType, &TransportationTypeDomain); err != nil {
		return nil, err
	}
	return &TransportationTypeDomain, nil
}

func (r *TransportationTypeRepo) Update(ctx context.Context, id domain.TransportationTypeId, TransportationTypeDomain domain.TransportationType) (*domain.TransportationType, error) {
	TransportationType := types.TransportationType{}
	if err := mapper.ConvertTypes(TransportationTypeDomain, &TransportationType); err != nil {
		return nil, err
	}
	if err := UpdateRecord(r.db, id, TransportationType); err != nil {
		return nil, err
	}

	domain, err := r.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return domain, nil
}

func (r *TransportationTypeRepo) GetByID(ctx context.Context, id domain.TransportationTypeId) (*domain.TransportationType, error) {
	TransportationType, err := GetRecordByID[types.TransportationType](r.db, id, nil)
	if err != nil {
		return nil, err
	}
	if TransportationType == nil {
		return nil, gorm.ErrRecordNotFound
	}

	TransportationTypeDomain := domain.TransportationType{}
	if err := mapper.ConvertTypes(*TransportationType, &TransportationTypeDomain); err != nil {
		return nil, err
	}

	return &TransportationTypeDomain, nil
}

func (r *TransportationTypeRepo) Delete(ctx context.Context, id domain.TransportationTypeId) error {
	return DeleteRecordByID[types.TransportationType](r.db, id)
}

func (r *TransportationTypeRepo) Get(ctx context.Context, request *commonDomain.RepositoryRequest) ([]domain.TransportationType, error) {
	companies, err := GetRecords[types.TransportationType](r.db, request)
	if err != nil {
		return []domain.TransportationType{}, err
	}

	TransportationTypeDomains := []domain.TransportationType{}

	mapper.ConvertTypes(companies, &TransportationTypeDomains)
	return TransportationTypeDomains, nil
}
