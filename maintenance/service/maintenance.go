package service

import (
	"context"
	"github.com/huy-quang-vmo/common_middleware/maintenance/entity"
	"github.com/huy-quang-vmo/common_middleware/maintenance/util"
	"slices"
)

type IMaintenanceRepo interface {
	UpdateServiceManagement(ctx context.Context, serviceManagement *entity.ServiceManagement) error
	GetServiceManagement(ctx context.Context) (entity.ServiceManagement, error)
	GetServiceStatus(ctx context.Context) (entity.ServiceStatus, error)
}

type MaintenanceService struct {
	repo IMaintenanceRepo
}

func NewMaintenanceService(repo IMaintenanceRepo) *MaintenanceService {
	return &MaintenanceService{repo: repo}
}

func (s *MaintenanceService) IsMaintenance() (bool, error) {
	status, err := s.repo.GetServiceStatus(context.Background())
	if err != nil {
		return false, err
	}

	return status == entity.StatusMaintenance, nil
}

func (s *MaintenanceService) UpdateStatus(status entity.ServiceStatus) error {
	// get maintenance service
	serviceManagement, err := s.repo.GetServiceManagement(context.Background())
	if err != nil {
		return err
	}

	if !slices.Contains([]entity.ServiceStatus{entity.StatusActive, entity.StatusInactive, entity.StatusMaintenance}, status) {
		return util.ErrInvalidServiceStatus
	}
	serviceManagement.Status = status
	return s.repo.UpdateServiceManagement(context.Background(), &serviceManagement)
}
