package client

import (
	"context"
	"food-truck-api/package/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreateClient(client *entities.Client) (*entities.Client, error)
}

type repository struct {
	Collection *mongo.Collection
}

func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}

func (r *repository) CreateClient(client *entities.Client) (*entities.Client, error) {
	client.ID = primitive.NewObjectID()
	client.UpdatedAt = time.Now()
	client.CreatedAt = time.Now()

	_, errr := r.Collection.InsertOne(context.Background(), client)

	if errr != nil {
		return nil, errr
	}

	return client, errr
}
