package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func connectToDB(host, port, dbname, username, password string) (*mongo.Database, error) {
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
