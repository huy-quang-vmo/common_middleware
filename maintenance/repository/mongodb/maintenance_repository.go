package mongodb

import (
	"context"
	"github.com/huy-quang-vmo/common_middleware/maintenance/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ServiceManagementCollection = "service_management"
)

type MaintenanceRepository struct {
	db *mongo.Database
}

func (m MaintenanceRepository) UpdateServiceManagement(ctx context.Context, serviceManagement *entity.ServiceManagement) error {
	// update service management collection
	collection := m.db.Collection(ServiceManagementCollection)
	update := bson.M{
		"$set": serviceManagement,
	}
	_, err := collection.UpdateByID(ctx, serviceManagement.ID, update)
	if err != nil {
		return err
	}

	return nil
}

func (m MaintenanceRepository) GetServiceManagement(ctx context.Context) (entity.ServiceManagement, error) {
	// get first record in service management collection
	collection := m.db.Collection(ServiceManagementCollection)
	filter := bson.D{}
	var serviceManagement entity.ServiceManagement
	err := collection.FindOne(ctx, filter).Decode(&serviceManagement)
	if err != nil {
		return serviceManagement, err
	}

	return serviceManagement, nil
}

func (m MaintenanceRepository) GetServiceStatus(ctx context.Context) (entity.ServiceStatus, error) {
	// get first record in service management collection
	collection := m.db.Collection(ServiceManagementCollection)
	filter := bson.D{}
	var serviceManagement entity.ServiceManagement
	err := collection.FindOne(ctx, filter).Decode(&serviceManagement)
	if err != nil {
		return serviceManagement.Status, err
	}

	return serviceManagement.Status, nil
}

func NewMaintenanceRepository(db *mongo.Database) *MaintenanceRepository {
	return &MaintenanceRepository{
		db: db,
	}
}
