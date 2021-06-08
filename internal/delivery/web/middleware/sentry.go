package middleware

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
)

func Sentry(Dsn string) echo.MiddlewareFunc {
	err := sentry.Init(sentry.ClientOptions{Dsn: Dsn})
	if err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	option := sentryecho.Options{Repanic: true}
	return sentryecho.New(option)
}
