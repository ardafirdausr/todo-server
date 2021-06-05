package mongo

import (
	"context"
	"log"

	"github.com/ardafirdausr/todo-server/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	DB *mongo.Database
}

func NewUserRepository(DB *mongo.Database) *UserRepository {
	return &UserRepository{DB: DB}
}

func (repo UserRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	res := repo.DB.Collection("users").FindOne(context.TODO(), bson.M{"email": email})
	if res.Err() == mongo.ErrNoDocuments {
		log.Println(res.Err())
		err := entity.NewErrNotFound("User not found", res.Err())
		return nil, err
	}

	if err := res.Decode(&user); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &user, nil
}

func (repo UserRepository) Create(t entity.CreateUserParam) (*entity.User, error) {
	r, err := repo.DB.Collection("users").InsertOne(context.TODO(), t)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	user := &entity.User{
		ID:       r.InsertedID.(primitive.ObjectID),
		Email:    t.Email,
		Name:     t.Name,
		ImageUrl: t.ImageUrl,
	}
	return user, nil
}
