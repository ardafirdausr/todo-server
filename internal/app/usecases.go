package app

import (
	"github.com/ardafirdausr/todo-server/internal"
	"github.com/ardafirdausr/todo-server/internal/usecase"
)

type Usecases struct {
	TodoUsecase internal.TodoUsecase
	AuthUsecase internal.AuthUsecase
}

func NewUsecases(repositories *Repositories) *Usecases {
	return &Usecases{
		TodoUsecase: usecase.NewTodoUsecase(repositories.TodoRepository),
		AuthUsecase: usecase.NewAuthUsecase(repositories.UserRepository),
	}
}
