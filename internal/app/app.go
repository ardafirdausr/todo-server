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
		log.Fatalf("Failed to load .env  file \n%v", err)
	}

	app.Repositories, err = NewRepositories()
	if err != nil {
		log.Fatalf("Failed to initiate repositories\n%v", err)
	}

	app.Usecases = NewUsecases(app.Repositories)
	return app, nil
}
