package usecase

import (
	"github.com/ardafirdausr/todo-server/internal"
	"github.com/ardafirdausr/todo-server/internal/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoUsecase struct {
	todoRepository internal.TodoRepository
}

func NewTodoUsecase(todoRepository internal.TodoRepository) *TodoUsecase {
	return &TodoUsecase{todoRepository: todoRepository}
}

func (service TodoUsecase) GetAllUserTodos(ID primitive.ObjectID) ([]*entity.Todo, error) {
	return service.todoRepository.GetTodosByUserID(ID)
}

func (service TodoUsecase) CreateTodo(t entity.CreateTodoParam) (*entity.Todo, error) {
	return service.todoRepository.Create(t)
}

func (service TodoUsecase) UpdateTodo(ID primitive.ObjectID, t entity.UpdateTodoParam) (*entity.Todo, error) {
	return service.todoRepository.UpdateById(ID, t)
}

func (service TodoUsecase) DeleteTodo(ID primitive.ObjectID) (bool, error) {
	return service.todoRepository.DeleteById(ID)
}
