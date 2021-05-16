package mongo

import (
	"context"
	"log"

	"github.com/ardafirdausr/todo-server/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TodoRepository struct {
	DB *mongo.Database
}

func NewTodoRepository(DB *mongo.Database) *TodoRepository {
	return &TodoRepository{DB: DB}
}

func (repo TodoRepository) GetTodosByUserID(ID primitive.ObjectID) ([]entity.Todo, error) {
	ctx := context.TODO()
	csr, err := repo.DB.Collection("todos").Find(ctx, bson.M{"userId": ID})
	if err != nil {
		return nil, err
	}
	defer csr.Close(ctx)

	todos := make([]entity.Todo, 0)
	for csr.Next(ctx) {
		var todo entity.Todo
		err := csr.Decode(&todo)
		if err != nil {
			log.Fatal(err.Error())
		}

		todos = append(todos, todo)
	}

	return todos, nil
}

func (repo TodoRepository) Create(t entity.CreateTodoParam) (*entity.Todo, error) {
	ctx := context.TODO()
	r, err := repo.DB.Collection("todos").InsertOne(ctx, t)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	todo := &entity.Todo{
		ID:        r.InsertedID.(primitive.ObjectID),
		Task:      t.Task,
		Completed: t.Completed,
		UserID:    t.UserID,
	}
	return todo, nil
}

func (repo TodoRepository) UpdateById(ID primitive.ObjectID, t entity.UpdateTodoParam) (*entity.Todo, error) {
	ctx := context.TODO()
	_, err := repo.DB.Collection("todos").UpdateByID(ctx, ID, bson.M{"$set": t})
	if err != nil {
		return nil, err
	}

	todo := &entity.Todo{
		ID:        ID,
		Task:      t.Task,
		Completed: t.Completed,
	}
	return todo, nil
}

func (repo TodoRepository) DeleteById(ID primitive.ObjectID) (bool, error) {
	ctx := context.TODO()
	r, err := repo.DB.Collection("todos").DeleteOne(ctx, bson.M{"_id": ID})
	if err != nil {
		return false, err
	}

	return r.DeletedCount > 0, nil
}
