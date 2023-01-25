package company

import (
	"food-truck-api/package/company/contract"
	"food-truck-api/package/entities"
)

type service struct {
	repository contract.Service
}

//NewService is used to create a single instance of the service
func NewService(r contract.Repository) contract.Service {
	return &service{
		repository: r,
	}
}

func (s *service) Register(company *contract.RegisterRequest) (*entities.Company, error) {
	return s.repository.Register(company)
}
