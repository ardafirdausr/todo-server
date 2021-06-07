package main

import (
	"log"

	"github.com/ardafirdausr/todo-server/internal/app"
	"github.com/ardafirdausr/todo-server/internal/delivery/web"
)

func main() {
	app, err := app.New()
	if err != nil {
		log.Fatalf("Failed initiate the app\n%v", err)
	}

	web.Start(app)
}
