package internal

import (
	"github.com/ardafirdausr/todo-server/internal/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthUsecase interface {
	SSO(token string, authenticator SSOAuthenticator) (*entity.User, error)
	GenerateAuthToken(user entity.User, tokenizer Tokenizer) (string, error)
}

type TodoUsecase interface {
	GetAllUserTodos(primitive.ObjectID) ([]*entity.Todo, error)
	CreateTodo(entity.CreateTodoParam) (*entity.Todo, error)
	UpdateTodo(primitive.ObjectID, entity.UpdateTodoParam) (*entity.Todo, error)
	DeleteTodo(primitive.ObjectID) (bool, error)
}
