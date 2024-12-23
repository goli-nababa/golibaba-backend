package http

import (
	"context"
	"transportation/api/http/services"
	"transportation/app"
	"transportation/config"
)

func companyServiceGetter(appContainer app.App, cfg config.ServerConfig) ServiceGetter[*services.CompanyService] {
	return func(ctx context.Context) *services.CompanyService {
		return services.NewCompanyService(appContainer.CompanyService(ctx))

	}
}
