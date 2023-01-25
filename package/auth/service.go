package auth

import "food-truck-api/package/entities"

type Service interface {
	HashPassword(password string) (string, error)
	CheckPassword(hash string, password string) bool
	CreateToken(company *entities.Company) (string, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) HashPassword(password string) (string, error) {
	return s.repository.HashPassword(password)
}

func (s *service) CheckPassword(hash string, password string) bool {
	return s.repository.CheckPassword(hash, password)
}

func (s *service) CreateToken(company *entities.Company) (string, error) {
	return s.repository.CreateToken(company)
}
