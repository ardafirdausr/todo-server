package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

var DB *mongo.Database

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env  file \n%v", err)
	}

	DB, err = connectToDB(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"))
	if err != nil {
		log.Fatalf("Failed to connect to the database\n%v", err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(Auth)

	e.GET("/todos", GetAllTodos)
	e.POST("/todos", CreateTodo)
	e.PUT("/todos/:id", UpdateTodo)
	e.DELETE("/todos/:id", DeleteTodo)

	e.Static("/", "web")

	fmt.Println("Serving app on port " + os.Getenv("PORT"))
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
