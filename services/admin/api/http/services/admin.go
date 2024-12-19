package services

import adminService "admin/internal/admin/port"

type AdminService struct {
	svc adminService.Service
}

func NewAdminService(svc adminService.Service) *AdminService {
	return &AdminService{
		svc: svc,
	}
}
