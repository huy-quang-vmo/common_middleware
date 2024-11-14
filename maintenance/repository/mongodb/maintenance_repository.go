package mongodb

import (
	"context"
	"errors"
	"github.com/huy-quang-vmo/common_middleware/maintenance/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	ServiceManagementCollection = "service_management"
)

type MaintenanceRepository struct {
	db *mongo.Database
}

func (m MaintenanceRepository) UpdateServiceManagement(ctx context.Context, serviceManagement *entity.ServiceManagement) error {
	collection := m.db.Collection(ServiceManagementCollection)
	filter := bson.M{"_id": serviceManagement.ID}
	update := bson.M{
		"$set": serviceManagement,
	}
	opts := options.Update().SetUpsert(true)
	_, err := collection.UpdateOne(ctx, filter, update, opts)
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
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return serviceManagement, err
	}

	if errors.Is(err, mongo.ErrNoDocuments) {
		return entity.ServiceManagement{
			ID:     primitive.NewObjectID(),
			Status: entity.StatusActive,
		}, nil
	}

	return serviceManagement, nil
}

func (m MaintenanceRepository) GetServiceStatus(ctx context.Context) (entity.ServiceStatus, error) {
	// get first record in service management collection
	collection := m.db.Collection(ServiceManagementCollection)
	filter := bson.D{}
	var serviceManagement entity.ServiceManagement
	err := collection.FindOne(ctx, filter).Decode(&serviceManagement)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return serviceManagement.Status, err
	}

	if errors.Is(err, mongo.ErrNoDocuments) {
		return entity.StatusActive, nil
	}

	return serviceManagement.Status, nil
}

func NewMaintenanceRepository(db *mongo.Database) *MaintenanceRepository {
	return &MaintenanceRepository{
		db: db,
	}
}
