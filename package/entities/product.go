package entities

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `json:"id"  bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description,omitempty"`
	CompanyId   string             `json:"companyId" bson:"companyId"`

	Images      []string             `json:"images" bson:"images"`
	Ingredients []primitive.ObjectID `json:"ingredients" bson:"ingredients,omitempty"`
	Category    primitive.ObjectID   `json:"category" bson:"category"`

	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updatedAt"`
}
