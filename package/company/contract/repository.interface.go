package contract

import "food-truck-api/package/entities"

type Repository interface {
	Register(company *RegisterRequest) (*entities.Company, error)
}
