package contract

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateTokenRequest struct {
	ID    primitive.ObjectID
	Name  string
	Email string
}
