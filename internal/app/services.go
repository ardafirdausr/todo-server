package app

import (
	"github.com/ardafirdausr/todo-server/internal"
	"github.com/ardafirdausr/todo-server/internal/service"
)

type Services struct {
	TodoService internal.TodoService
	AuthService internal.AuthService
}

func NewServices(repositories *Repositories) *Services {
	return &Services{
		TodoService: service.NewTodoService(repositories.TodoRepository),
		AuthService: service.NewAuthService(repositories.UserRepository),
	}
}
