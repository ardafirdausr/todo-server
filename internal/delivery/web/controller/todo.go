package controller

import (
	"net/http"

	"github.com/ardafirdausr/todo-server/internal/app"
	"github.com/ardafirdausr/todo-server/internal/entity"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoController struct {
	services *app.Usecases
}

func NewTodoController(services *app.Usecases) *TodoController {
	return &TodoController{services: services}
}

func (ctrl TodoController) GetAllTodos(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*entity.JWTPayload)
	userId := claims.ID
	todos, err := ctrl.services.TodoUsecase.GetAllUserTodos(userId)
	if err != nil {
		c.Logger().Error(err.Error())
		return err
	}

	payload := echo.Map{
		"message": http.StatusText(http.StatusOK),
		"data":    todos,
	}
	return c.JSON(http.StatusOK, payload)
}

func (ctrl TodoController) CreateTodo(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*entity.JWTPayload)
	userId := claims.ID

	todo := entity.CreateTodoParam{
		UserID:    userId,
		Completed: false,
	}
	if err := c.Bind(&todo); err != nil {
		c.Logger().Error(err.Error())
		return echo.ErrInternalServerError
	}

	if err := c.Validate(&todo); err != nil {
		c.Logger().Error(err.Error())
		return err
	}

	newTodo, err := ctrl.services.TodoUsecase.CreateTodo(todo)
	if err != nil {
		c.Logger().Error(err.Error())
		return err
	}

	payload := echo.Map{
		"message": http.StatusText(http.StatusOK),
		"data":    newTodo,
	}
	return c.JSON(http.StatusCreated, payload)
}

func (ctrl TodoController) UpdateTodo(c echo.Context) error {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.Logger().Error(err.Error())
		payload := echo.Map{"message": "Invalid ID"}
		return c.JSON(http.StatusBadRequest, payload)
	}

	todo, err := ctrl.services.TodoUsecase.GetTodo(objectId)
	if err != nil {
		c.Logger().Error(err.Error())
		return err
	}

	if err := c.Bind(todo); err != nil {
		c.Logger().Error(err.Error())
		return echo.ErrInternalServerError
	}

	updateData := entity.UpdateTodoParam{
		Task:      todo.Task,
		Completed: todo.Completed,
	}

	if err := c.Validate(todo); err != nil {
		c.Logger().Error(err.Error())
		return err
	}

	isUpdated, err := ctrl.services.TodoUsecase.UpdateTodo(objectId, updateData)
	if err != nil {
		c.Logger().Error(err.Error())
		return err
	}

	if !isUpdated {
		return echo.ErrInternalServerError
	}

	payload := echo.Map{
		"message": http.StatusText(http.StatusOK),
		"data":    todo,
	}
	return c.JSON(http.StatusOK, payload)
}

func (ctrl TodoController) DeleteTodo(c echo.Context) error {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.Logger().Error(err.Error())
		payload := echo.Map{"message": "Invalid ID"}
		return c.JSON(http.StatusBadRequest, payload)
	}

	isDeleted, err := ctrl.services.TodoUsecase.DeleteTodo(objectId)
	if err != nil {
		c.Logger().Error(err.Error())
		return err
	}

	if !isDeleted {
		return echo.ErrInternalServerError
	}

	return c.NoContent(http.StatusNoContent)
}
