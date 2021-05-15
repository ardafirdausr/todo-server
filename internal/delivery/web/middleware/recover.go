package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func Recover() echo.MiddlewareFunc {
	config := middleware.RecoverConfig{
		LogLevel: log.ERROR,
	}
	return middleware.RecoverWithConfig(config)
}
