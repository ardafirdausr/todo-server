package internal

import (
	"github.com/ardafirdausr/todo-server/internal/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoRepository interface {
	GetTodosByUserID(primitive.ObjectID) ([]*entity.Todo, error)
	Create(entity.CreateTodoParam) (*entity.Todo, error)
	UpdateById(primitive.ObjectID, entity.UpdateTodoParam) (*entity.Todo, error)
	DeleteById(primitive.ObjectID) (bool, error)
}

type UserRepository interface {
	GetUserByEmail(string) (*entity.User, error)
	Create(entity.CreateUserParam) (*entity.User, error)
}
