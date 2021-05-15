package web

import (
	"github.com/ardafirdausr/todo-server/internal/app"
	"github.com/ardafirdausr/todo-server/internal/delivery/web/controller"
	"github.com/ardafirdausr/todo-server/internal/delivery/web/middleware"
	"github.com/ardafirdausr/todo-server/internal/delivery/web/server"
)

func Start(app *app.TodoApp) {
	web := server.New()

	todoController := controller.NewTodoController(app.Services)
	todoGroup := web.Group("/todos")
	todoGroup.Use(middleware.JWT())
	todoGroup.GET("", todoController.GetAllTodos)
	todoGroup.POST("", todoController.CreateTodo)
	todoGroup.PUT("/:id", todoController.UpdateTodo)
	todoGroup.DELETE("/:id", todoController.DeleteTodo)

	authController := controller.NewAuthController(app.Services)
	authGroup := web.Group("/auth")
	authGroup.POST("/login", authController.Login)

	server.Start(web)
}
