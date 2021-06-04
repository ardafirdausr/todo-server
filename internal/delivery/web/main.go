package web

import (
	"os"

	"github.com/ardafirdausr/todo-server/internal/app"
	"github.com/ardafirdausr/todo-server/internal/delivery/web/controller"
	"github.com/ardafirdausr/todo-server/internal/delivery/web/middleware"
	"github.com/ardafirdausr/todo-server/internal/delivery/web/server"
)

func Start(app *app.TodoApp) {
	web := server.New()

	JWTSecretKey := os.Getenv("JWT_SECRET_KEY")
	JWTmiddleware := middleware.JWT(JWTSecretKey)

	todoController := controller.NewTodoController(app.Usecases)
	todoGroup := web.Group("/todos")
	todoGroup.Use(JWTmiddleware)
	todoGroup.GET("", todoController.GetAllTodos)
	todoGroup.POST("", todoController.CreateTodo)
	todoGroup.PUT("/:id", todoController.UpdateTodo)
	todoGroup.DELETE("/:id", todoController.DeleteTodo)

	authController := controller.NewAuthController(app.Usecases)
	authGroup := web.Group("/auth")
	authGroup.POST("/login", authController.Login)

	server.Start(web)
}
