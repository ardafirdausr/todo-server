package app

import (
	"log"

	"github.com/joho/godotenv"
)

type TodoApp struct {
	Repositories *Repositories
	Services     *Services
}

func New() (*TodoApp, error) {
	app := new(TodoApp)

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env  file \n%v", err)
	}

	app.Repositories, err = NewRepositories()
	if err != nil {
		log.Fatalf("Failed to initiate repositories\n%v", err)
	}

	app.Services = NewServices(app.Repositories)
	return app, nil
}
