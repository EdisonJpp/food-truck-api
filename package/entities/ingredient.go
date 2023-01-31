package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ingredient struct {
	ID        primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Name      string             `json:"name" bson:"name"`
	CompanyId string             `json:"companyId" bson:"companyId"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
