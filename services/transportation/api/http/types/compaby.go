package types

import (
	"transportation/internal/company/domain"
	transportationDomain "transportation/internal/transportation_type/domain"
)

type CreateCompanyRequest struct {
	Name                 string                                    `json:"name"`
	OwnerId              uint                                      `json:"owner_id"`
	TransportationTypeId transportationDomain.TransportationTypeId `json:"transportation_type_id"`
}
type CompanyResponse struct {
	Id                 domain.CompanyId                        `json:"id"`
	Name               string                                  `json:"name"`
	OwnerId            uint                                    `json:"owner_id"`
	TransportationType transportationDomain.TransportationType `json:"transportation_type"`
}

type UpdateCompanyRequest struct {
	Name                 string                                    `json:"name"`
	TransportationTypeId transportationDomain.TransportationTypeId `json:"transportation_type_id"`
}

type FilterCompaniesRequest struct {
	OwnerId              uint                                      `json:"owner_id"`
	TransportationTypeId transportationDomain.TransportationTypeId `json:"transportation_type_id"`
}
