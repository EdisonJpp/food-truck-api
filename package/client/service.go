package client

import "food-truck-api/package/entities"

type Service interface {
	CreateClient(client *entities.Client) (*entities.Client, error)
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
