package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	e.GET("/todos", GetAllTodos)
	e.POST("/todos", CreateTodo)
	e.PUT("/todos/:id", UpdateTodo)
	e.DELETE("/todos/:id", DeleteTodo)

	e.Start(":8080")
}
