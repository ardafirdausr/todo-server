package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(e echo.Context) error {
		return e.String(http.StatusAccepted, "Hello World")
	})

	e.Start(":8080")
}
