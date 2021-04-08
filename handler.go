package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllTodos(c echo.Context) error {
	todos, err := FindAllTodos()
	if err != nil {
		log.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, todos)
}

func CreateTodo(c echo.Context) error {
	todo := Todo{}
	if err := c.Bind(&todo); err != nil {
		return err
	}

	newTodo, err := todo.Save()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, newTodo)
}

func UpdateTodo(c echo.Context) error {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		payload := echo.Map{"message": "Invalid ID"}
		return c.JSON(http.StatusOK, payload)
	}

	todo := Todo{ID: objectId}
	if err := c.Bind(&todo); err != nil {
		return err
	}

	updatedTodo, err := todo.Update()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, updatedTodo)
}

func DeleteTodo(c echo.Context) error {
	id := c.Param("id")
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		payload := echo.Map{"message": "Invalid ID"}
		return c.JSON(http.StatusOK, payload)
	}

	todo := Todo{ID: objectId}
	if err := c.Bind(&todo); err != nil {
		return err
	}

	isDeleted, err := todo.Delete()
	if err != nil {
		return err
	}

	if !isDeleted {
		payload := echo.Map{"message": "Failed to delete todo"}
		return c.JSON(http.StatusInternalServerError, payload)
	}

	return c.JSON(http.StatusNoContent, nil)
}
