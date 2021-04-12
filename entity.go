package main

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type key int

const (
	userIdKey key = iota
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	ImageUrl string             `json:"image_url" bson:"imageUrl"`
}

func (u *User) Login() error {
	res := DB.Collection("users").FindOne(context.TODO(), bson.M{"email": u.Email})
	if err := res.Decode(u); err == nil {
		return nil
	}

	r, err := DB.Collection("users").InsertOne(context.TODO(), u)
	if err != nil {
		log.Println(err)
		return err
	}

	u.ID = r.InsertedID.(primitive.ObjectID)
	return nil
}

type Todo struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Task      string             `json:"task" bson:"task"`
	Completed bool               `json:"completed" bson:"completed"`
	UserID    primitive.ObjectID `json:"user_id" bson:"userId,omitempty"`
}

func (t *Todo) Save() (*Todo, error) {
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

func GetAllUserTodos(userId primitive.ObjectID) ([]Todo, error) {
	ctx := context.TODO()
	csr, err := DB.Collection("todos").Find(ctx, bson.M{"userId": userId})
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
