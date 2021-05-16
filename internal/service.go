package internal

import (
	"github.com/ardafirdausr/todo-server/internal/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/api/idtoken"
)

type AuthService interface {
	Register(entity.CreateUserParam) (*entity.User, error)
	GetUserByEmail(string) (*entity.User, error)
	GetGoogleAuthPayload(entity.GoogleAuth) (*idtoken.Payload, error)
	GenerateJWTToken(entity.JWTPayload) (string, error)
}

type TodoService interface {
	GetAllUserTodos(primitive.ObjectID) ([]entity.Todo, error)
	CreateTodo(entity.CreateTodoParam) (*entity.Todo, error)
	UpdateTodo(primitive.ObjectID, entity.UpdateTodoParam) (*entity.Todo, error)
	DeleteTodo(primitive.ObjectID) (bool, error)
}
