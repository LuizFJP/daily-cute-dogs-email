package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

type Email struct {
	ID    primitive.ObjectID `bson:"_id, omitempty"`
	Email string             `bson:"email"`
}

