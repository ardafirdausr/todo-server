package service

import (
	"github.com/ardafirdausr/todo-server/internal"
	"github.com/ardafirdausr/todo-server/internal/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoService struct {
	todoRepository internal.TodoRepository
}

func NewTodoService(todoRepository internal.TodoRepository) *TodoService {
	return &TodoService{todoRepository: todoRepository}
}

func (service TodoService) GetAllUserTodos(ID primitive.ObjectID) ([]entity.Todo, error) {
	return service.todoRepository.GetTodosByUserID(ID)
}

func (service TodoService) CreateTodo(t entity.CreateTodoParam) (*entity.Todo, error) {
	return service.todoRepository.Create(t)
}

func (service TodoService) UpdateTodo(ID primitive.ObjectID, t entity.UpdateTodoParam) (*entity.Todo, error) {
	return service.todoRepository.UpdateById(ID, t)
}

func (service TodoService) DeleteTodo(ID primitive.ObjectID) (bool, error) {
	return service.todoRepository.DeleteById(ID)
}
