package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Todo struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Task      string             `json:"task" bson:"task"`
	Completed bool               `json:"completed" bson:"completed"`
}

func (t *Todo) Save() (*Todo, error) {
	t.ID = primitive.NewObjectID()

	ctx := context.TODO()
	r, err := DB.Collection("todos").InsertOne(ctx, t)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	t.ID = r.InsertedID.(primitive.ObjectID)
	return t, nil
}

func (t *Todo) Update() (*Todo, error) {
	ctx := context.TODO()
	_, err := DB.Collection("todos").UpdateByID(ctx, t.ID, bson.M{"$set": t})
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (t *Todo) Delete() (bool, error) {
	ctx := context.TODO()
	r, err := DB.Collection("todos").DeleteOne(ctx, bson.M{"_id": t.ID})
	if err != nil {
		return false, err
	}

	return r.DeletedCount > 0, nil
}

func FindAllTodos() ([]Todo, error) {
	ctx := context.TODO()
	csr, err := DB.Collection("todos").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer csr.Close(ctx)

	todos := make([]Todo, 0)
	for csr.Next(ctx) {
		var row Todo
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		todos = append(todos, row)
	}

	return todos, nil
}
