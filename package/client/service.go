package client

import (
	"food-truck-api/package/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	CreateClient(client *entities.Client) (*entities.Client, error)
	GetClientByEmail(email string, companyId primitive.ObjectID) (*entities.Client, error)
	IsEmailExists(email string, companyId primitive.ObjectID) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) CreateClient(client *entities.Client) (*entities.Client, error) {
	return s.repository.CreateClient(client)
}

func (s *service) IsEmailExists(email string, companyId primitive.ObjectID) (bool, error) {
	return s.repository.IsEmailExists(email, companyId)
}

func (s *service) GetClientByEmail(email string, companyId primitive.ObjectID) (*entities.Client, error) {
	return s.repository.GetClientByEmail(email, companyId)
}
