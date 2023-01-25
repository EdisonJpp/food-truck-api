package product

import (
	"context"
	"food-truck-api/package/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GetProducts() (*[]entities.Product, error)
}

type repository struct {
	Collection *mongo.Collection
}

func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}

func (r *repository) GetProducts() (*[]entities.Product, error) {
	var allProducts []entities.Product

	cursor, err := r.Collection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var product entities.Product
		_ = cursor.Decode(&product)
		allProducts = append(allProducts, product)
	}

	return &allProducts, nil
}
