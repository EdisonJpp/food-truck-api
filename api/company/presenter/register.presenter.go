package presenter

import (
	"food-truck-api/package/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RegisterPresentation struct {
	ID          primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Name        string             `json:"name"`
	Email       string             `json:"email"`
	AccessToken string             `json:"accessToken"`
}

func RegisterPresent(register *entities.Company, accessToken string) RegisterPresentation {
	return RegisterPresentation{
		ID:          register.ID,
		Name:        register.Name,
		Email:       register.Email,
		AccessToken: accessToken,
	}
}
