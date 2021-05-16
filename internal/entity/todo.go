package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Task      string             `json:"task" bson:"task"`
	Completed bool               `json:"completed" bson:"completed"`
	UserID    primitive.ObjectID `json:"userId" bson:"userId,omitempty"`
}

type CreateTodoParam struct {
	Task      string             `json:"task" bson:"task" validate:"required"`
	Completed bool               `json:"completed" bson:"completed" validate:"required"`
	UserID    primitive.ObjectID `json:"userId" bson:"userId,omitempty" validate:"required"`
}

type UpdateTodoParam struct {
	Task      string `json:"task" bson:"task"`
	Completed bool   `json:"completed" bson:"completed"`
}
