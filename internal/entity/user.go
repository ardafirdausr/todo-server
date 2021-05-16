package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	ImageUrl string             `json:"imageUrl" bson:"imageUrl"`
}

type CreateUserParam struct {
	Name     string `json:"name" bson:"name" validate:"required"`
	Email    string `json:"email" bson:"email" validate:"required"`
	ImageUrl string `json:"imageUrl" bson:"imageUrl" validate:"required"`
}
