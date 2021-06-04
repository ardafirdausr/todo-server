package usecase

import (
	"testing"

	"github.com/ardafirdausr/todo-server/internal/entity"
	"github.com/ardafirdausr/todo-server/internal/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetAllUserTodos_Success(t *testing.T) {
	var dummyUserID primitive.ObjectID = primitive.NewObjectID()
	var dummyTodos []*entity.Todo = []*entity.Todo{
		{
			ID:        primitive.NewObjectID(),
			Task:      "Task 1",
			Completed: true,
			UserID:    dummyUserID,
		},
		{
			ID:        primitive.NewObjectID(),
			Task:      "Task 2",
			Completed: false,
			UserID:    dummyUserID,
		},
	}

	mockTodoRepository := new(mocks.TodoRepository)
	mockTodoRepository.On("GetTodosByUserID", dummyUserID).Return(dummyTodos, nil)

	todoUsecase := NewTodoUsecase(mockTodoRepository)
	todos, err := todoUsecase.GetAllUserTodos(dummyUserID)
	assert.Equal(t, dummyTodos, todos)
	assert.Equal(t, len(dummyTodos), len(todos))
	assert.Nil(t, err)
}

func TestCreateTodo_Success(t *testing.T) {
	dummyUserID := primitive.NewObjectID()
	willCreateTodo := entity.CreateTodoParam{
		Task:      "Task 1",
		Completed: false,
		UserID:    dummyUserID,
	}
	expectedTodo := entity.Todo{
		ID:        primitive.NewObjectID(),
		Task:      "Task 1",
		Completed: false,
		UserID:    dummyUserID,
	}

	mockTodoRepository := new(mocks.TodoRepository)
	mockTodoRepository.On("Create", willCreateTodo).Return(&expectedTodo, nil)

	todoUsecase := NewTodoUsecase(mockTodoRepository)
	todo, err := todoUsecase.CreateTodo(willCreateTodo)
	assert.ObjectsAreEqualValues(expectedTodo, todo)
	assert.Nil(t, err)
}

func TestUpdateTodo_Success(t *testing.T) {
	dummyUserID := primitive.NewObjectID()
	willUpdateTodo := entity.Todo{
		ID:        primitive.NewObjectID(),
		Task:      "Task 1",
		Completed: false,
		UserID:    dummyUserID,
	}
	newTodoData := entity.UpdateTodoParam{
		Task:      "Finished Task",
		Completed: true,
	}
	expectedTodo := entity.Todo{
		ID:        willUpdateTodo.ID,
		Task:      "Finished Task",
		Completed: true,
		UserID:    dummyUserID,
	}

	mockTodoRepository := new(mocks.TodoRepository)
	mockTodoRepository.On("UpdateById", willUpdateTodo.ID, newTodoData).Return(&expectedTodo, nil)

	todoUsecase := NewTodoUsecase(mockTodoRepository)
	todo, err := todoUsecase.UpdateTodo(willUpdateTodo.ID, newTodoData)
	assert.ObjectsAreEqualValues(expectedTodo, todo)
	assert.Nil(t, err)
}

func TestDeleteTodo_Success(t *testing.T) {
	dummyUserID := primitive.NewObjectID()
	willDeleteTodo := entity.Todo{
		ID:        primitive.NewObjectID(),
		Task:      "Task 1",
		Completed: false,
		UserID:    dummyUserID,
	}

	mockTodoRepository := new(mocks.TodoRepository)
	mockTodoRepository.On("DeleteById", willDeleteTodo.ID).Return(true, nil)

	todoUsecase := NewTodoUsecase(mockTodoRepository)
	isDeleted, err := todoUsecase.DeleteTodo(willDeleteTodo.ID)
	assert.True(t, isDeleted)
	assert.Nil(t, err)
}
