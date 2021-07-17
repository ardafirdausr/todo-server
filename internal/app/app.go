package app

import (
	"log"

	"github.com/joho/godotenv"
)

type TodoApp struct {
	Repositories *Repositories
	Usecases     *Usecases
}

func New() (*TodoApp, error) {
	app := new(TodoApp)

	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)

	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning! The .env file is not found. Make sure all env variables are declared \n%v", err)
	}

	app.Repositories, err = NewRepositories()
	if err != nil {
		log.Fatalf("Failed to initiate repositories\n%v", err)
	}

	app.Usecases = NewUsecases(app.Repositories)
	return app, nil
}
