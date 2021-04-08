package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var lastId = 3
var DB = []Todo{
	{ID: 1, Task: "Buy Eggs"},
	{ID: 2, Task: "Buy Milk"},
}

func GetAllTodos(c echo.Context) error {
	todos := DB
	return c.JSON(http.StatusOK, todos)
}

func CreateTodo(c echo.Context) error {
	todo := Todo{}
	if err := c.Bind(&todo); err != nil {
		return err
	}

	todo.ID = lastId
	lastId++

	DB = append(DB, todo)
	return c.JSON(http.StatusCreated, todo)
}

func UpdateTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	newTodo := Todo{}
	newTodo.ID = id
	if err := c.Bind(&newTodo); err != nil {
		return err
	}

	var updateIndex = -1
	for idx, todo := range DB {
		if todo.ID == newTodo.ID {
			updateIndex = idx
		}
	}

	if updateIndex == -1 {
		return c.JSON(http.StatusNotFound, nil)
	}

	DB[updateIndex] = newTodo
	return c.JSON(http.StatusOK, newTodo)
}

func DeleteTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var deleteIndex = -1
	for index, todo := range DB {
		if todo.ID == id {
			deleteIndex = index
		}
	}

	if deleteIndex == -1 {
		return c.JSON(http.StatusNotFound, nil)
	}

	DB = append(DB[:deleteIndex], DB[deleteIndex+1:]...)
	return c.JSON(http.StatusNoContent, nil)
}
