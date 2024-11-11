package service

import (
	"github.com/huy-quang-vmo/common_middleware/maintenance/entity"
	"github.com/huy-quang-vmo/common_middleware/maintenance/util"
	"slices"
)

type IMaintenanceRepo interface {
	GetServiceStatus() (entity.ServiceStatus, error)
	UpdateStatus(status entity.ServiceStatus) error
}

type MaintenanceService struct {
	repo IMaintenanceRepo
}

func NewMaintenanceService(repo IMaintenanceRepo) *MaintenanceService {
	return &MaintenanceService{repo: repo}
}

func (s *MaintenanceService) IsMaintenance() (bool, error) {
	status, err := s.repo.GetServiceStatus()
	if err != nil {
		return false, err
	}

	return status == entity.StatusMaintenance, nil
}

func (s *MaintenanceService) UpdateStatus(status entity.ServiceStatus) error {
	if !slices.Contains([]entity.ServiceStatus{entity.StatusActive, entity.StatusInactive, entity.StatusMaintenance}, status) {
		return util.ErrInvalidServiceStatus
	}
	return s.repo.UpdateStatus(status)
}
