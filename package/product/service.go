package product

import (
	"food-truck-api/package/entities"
)

//Service is an interface from which our api module can access our repository of all our models
type Service interface {
	GetProducts() (*[]entities.Product, error)
}

type service struct {
	repository Repository
}

//NewService is used to create a single instance of the service
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetProducts() (*[]entities.Product, error) {
	return s.repository.GetProducts()
}
