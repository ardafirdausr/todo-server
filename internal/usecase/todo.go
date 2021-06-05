package usecase

import (
	"log"

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
	todos, err := service.todoRepository.GetTodosByUserID(ID)
	if err != nil {
		log.Println(err.Error())
	}

	return todos, err
}

func (service TodoUsecase) CreateTodo(t entity.CreateTodoParam) (*entity.Todo, error) {
	todo, err := service.todoRepository.Create(t)
	if err != nil {
		log.Println(err.Error())
	}
	return todo, err
}

func (service TodoUsecase) UpdateTodo(ID primitive.ObjectID, t entity.UpdateTodoParam) (*entity.Todo, error) {
	todo, err := service.todoRepository.UpdateById(ID, t)
	if err != nil {
		log.Println(err.Error())
	}
	return todo, err
}

func (service TodoUsecase) DeleteTodo(ID primitive.ObjectID) (bool, error) {
	isDeleted, err := service.todoRepository.DeleteById(ID)
	if err != nil {
		log.Println(err.Error())
	}
	return isDeleted, err
}
