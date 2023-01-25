package contract

import "food-truck-api/package/entities"

type Repository interface {
	Register(company *RegisterRequest) (*entities.Company, error)
	GetCompanyByEmail(email string) (*entities.Company, error)
}
