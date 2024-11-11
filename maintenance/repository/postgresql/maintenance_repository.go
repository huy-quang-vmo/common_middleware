package postgresql

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/huy-quang-vmo/common_middleware/maintenance/entity"
)

type MaintenanceRepository struct {
	DB *pg.DB
}

func NewMaintenanceRepository(db *pg.DB) *MaintenanceRepository {
	return &MaintenanceRepository{DB: db}
}

func (i *MaintenanceRepository) GetServiceStatus(ctx context.Context) (entity.ServiceStatus, error) {
	var serviceManagement entity.ServiceManagement
	err := i.DB.ModelContext(ctx, &serviceManagement).First()
	if err != nil {
		return serviceManagement.Status, err
	}

	return serviceManagement.Status, nil
}

func (i *MaintenanceRepository) GetServiceManagement(ctx context.Context) (entity.ServiceManagement, error) {
	var serviceManagement entity.ServiceManagement
	err := i.DB.ModelContext(ctx, &serviceManagement).First()
	if err != nil {
		return serviceManagement, err
	}

	return serviceManagement, nil
}

func (i *MaintenanceRepository) UpdateServiceManagement(ctx context.Context, serviceManagement *entity.ServiceManagement) error {
	_, err := i.DB.ModelContext(ctx, serviceManagement).WherePK().Update()
	if err != nil {
		return err
	}

	return nil
}
