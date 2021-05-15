package middleware

import (
	"os"

	"github.com/ardafirdausr/todo-server/internal/entity"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func JWT() echo.MiddlewareFunc {
	JWTSecretKey := os.Getenv("JWT_SECRET_KEY")
	config := middleware.JWTConfig{
		Claims:     &entity.JWTPayload{},
		SigningKey: []byte(JWTSecretKey),
	}
	return middleware.JWTWithConfig(config)
}
