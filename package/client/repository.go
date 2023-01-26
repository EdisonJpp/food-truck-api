package client

import (
	"context"
	"food-truck-api/package/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreateClient(client *entities.Client) (*entities.Client, error)
	GetClientByEmail(email string) (*entities.Client, error)
	IsEmailExists(email string) bool
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

func (r *repository) IsEmailExists(email string) bool {
	filter := bson.D{{Key: "email", Value: email}}

	count, err := r.Collection.CountDocuments(context.TODO(), filter)

	if err != nil {
		panic(err)
	}

	return count > 0
}

func (r *repository) GetClientByEmail(email string) (*entities.Client, error) {
	client := new(entities.Client)

	filter := bson.D{{
		Key: "email", Value: email,
	}}

	err := r.Collection.FindOne(context.Background(), filter).Decode(&client)

	if err != nil {
		return nil, err
	}

	return client, nil
}
