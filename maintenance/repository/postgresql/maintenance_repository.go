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
