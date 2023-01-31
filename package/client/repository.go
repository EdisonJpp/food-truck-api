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
	GetClientByEmail(email string, companyId primitive.ObjectID) (*entities.Client, error)
	IsEmailExists(email string, companyId primitive.ObjectID) (bool, error)
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

func (r *repository) IsEmailExists(email string, companyId primitive.ObjectID) (bool, error) {
	filter := bson.D{{Key: "email", Value: email}, {Key: "companyId", Value: companyId}}

	count, err := r.Collection.CountDocuments(context.TODO(), filter)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *repository) GetClientByEmail(email string, companyId primitive.ObjectID) (*entities.Client, error) {
	client := new(entities.Client)
	filter := bson.D{{Key: "email", Value: email}, {Key: "companyId", Value: companyId}}

	err := r.Collection.FindOne(context.Background(), filter).Decode(&client)

	if err != nil {
		return nil, err
	}

	return client, nil
}
