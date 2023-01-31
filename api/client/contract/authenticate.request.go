package contract

import "go.mongodb.org/mongo-driver/bson/primitive"

type AuthenticateRequest struct {
	Email     string             `json:"email" validate:"email"`
	Name      string             `json:"password" validate:"gte=3,required"`
	CompanyId primitive.ObjectID `json:"companyId" validate:"required"`
}
