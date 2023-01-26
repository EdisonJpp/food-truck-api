package client

import "food-truck-api/package/entities"

type Service interface {
	CreateClient(client *entities.Client) (*entities.Client, error)
	GetClientByEmail(email string) (*entities.Client, error)
	IsEmailExists(email string) bool
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

func (s *service) IsEmailExists(email string) bool {
	return s.repository.IsEmailExists(email)
}

func (s *service) GetClientByEmail(email string) (*entities.Client, error) {
	return s.repository.GetClientByEmail(email)
}
