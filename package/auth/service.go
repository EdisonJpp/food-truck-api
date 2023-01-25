package auth

import (
	"food-truck-api/package/auth/contract"
)

type Service interface {
	HashPassword(password string) (string, error)
	CheckPassword(hash string, password string) bool
	CreateToken(payload *contract.CreateTokenRequest) (string, error)
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

func (s *service) CreateToken(payload *contract.CreateTokenRequest) (string, error) {
	return s.repository.CreateToken(payload)
}
