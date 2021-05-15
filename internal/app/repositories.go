package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ardafirdausr/todo-server/internal"
	mongoRepo "github.com/ardafirdausr/todo-server/internal/repository/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repositories struct {
	TodoRepository internal.TodoRepository
	UserRepository internal.UserRepository
}

func NewRepositories() (*Repositories, error) {
	DB, err := connectToMongoDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database\n%v", err)
		return nil, err
	}

	return &Repositories{
		TodoRepository: mongoRepo.NewTodoRepository(DB),
		UserRepository: mongoRepo.NewUserRepository(DB),
	}, nil
}

func connectToMongoDB() (*mongo.Database, error) {
	host := os.Getenv("MONGO_DB_HOST")
	port := os.Getenv("MONGO_DB_PORT")
	dbname := os.Getenv("MONGO_DB_NAME")
	username := os.Getenv("MONGO_DB_USERNAME")
	password := os.Getenv("MONGO_DB_PASSWORD")
	mongoDBURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=%s", username, password, host, port, dbname)

	clientOptions := options.Client().ApplyURI(mongoDBURI)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = client.Connect(ctx); err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return client.Database(dbname), nil
}
