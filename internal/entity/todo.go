package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Task      string             `json:"task" bson:"task"`
	Completed bool               `json:"completed" bson:"completed"`
	UserID    primitive.ObjectID `json:"userId" bson:"userId,omitempty"`
}

type CreateTodoParam struct {
	Task      string             `json:"task" bson:"task" validate:"required,min=1,max=50"`
	Completed bool               `json:"completed" bson:"completed"`
	UserID    primitive.ObjectID `bson:"userId,omitempty"`
}

type UpdateTodoParam struct {
	Task      string `json:"task,omitempty" bson:"task" validate:"min=1,max=50"`
	Completed bool   `json:"completed,omitempty" bson:"completed"`
}
