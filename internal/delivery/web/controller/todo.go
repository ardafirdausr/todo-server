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
		return err
	}

	return c.JSON(http.StatusOK, todos)
}

func (ctrl TodoController) CreateTodo(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*entity.JWTPayload)
	userId := claims.ID

	todo := entity.CreateTodoParam{UserID: userId}
	if err := c.Bind(&todo); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	newTodo, err := ctrl.services.TodoUsecase.CreateTodo(todo)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, newTodo)
}

func (ctrl TodoController) UpdateTodo(c echo.Context) error {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		payload := echo.Map{"message": "Invalid ID"}
		return c.JSON(http.StatusBadRequest, payload)
	}

	todo := entity.UpdateTodoParam{}
	if err := c.Bind(&todo); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	updatedTodo, err := ctrl.services.TodoUsecase.UpdateTodo(objectId, todo)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, updatedTodo)
}

func (ctrl TodoController) DeleteTodo(c echo.Context) error {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		payload := echo.Map{"message": "Invalid ID"}
		return c.JSON(http.StatusBadRequest, payload)
	}

	isDeleted, err := ctrl.services.TodoUsecase.DeleteTodo(objectId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if !isDeleted {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}
