package postgresql

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/huy-quang-vmo/common_middleware/maintenance/entity"
)

type INewsRepository struct {
	DB *pg.DB
}

func NewINewsRepository(db *pg.DB) *INewsRepository {
	return &INewsRepository{DB: db}
}

func (i *INewsRepository) GetServiceStatus(ctx context.Context) (entity.ServiceStatus, error) {
	var serviceManagement entity.ServiceManagement
	err := i.DB.ModelContext(ctx, &serviceManagement).First()
	if err != nil {
		return serviceManagement.Status, err
	}

	return serviceManagement.Status, nil
}

func (i *INewsRepository) GetServiceManagement(ctx context.Context) (entity.ServiceManagement, error) {
	var serviceManagement entity.ServiceManagement
	err := i.DB.ModelContext(ctx, &serviceManagement).First()
	if err != nil {
		return serviceManagement, err
	}

	return serviceManagement, nil
}

func (i *INewsRepository) UpdateServiceManagement(ctx context.Context, serviceManagement *entity.ServiceManagement) error {
	_, err := i.DB.ModelContext(ctx, serviceManagement).WherePK().Update()
	if err != nil {
		return err
	}

	return nil
}
