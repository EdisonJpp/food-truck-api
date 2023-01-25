package company

import (
	"context"
	"food-truck-api/package/company/contract"
	"food-truck-api/package/entities"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	Collection *mongo.Collection
}

func NewRepo(collection *mongo.Collection) contract.Repository {
	return &repository{
		Collection: collection,
	}
}

func (r *repository) Register(newCompany *contract.RegisterRequest) (*entities.Company, error) {
	company := new(entities.Company)

	company.ID = primitive.NewObjectID()
	company.CreatedAt = time.Now()
	company.UpdatedAt = time.Now()
	company.Email = newCompany.Email
	company.Name = newCompany.Name
	company.Password = newCompany.Password

	_, err := r.Collection.InsertOne(context.Background(), company)

	if err != nil {
		return nil, err
	}

	return company, nil
}
