package main

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
	"google.golang.org/api/idtoken"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		token := strings.Replace(authHeader, "Bearer ", "", -1)

		clientOption := idtoken.WithHTTPClient(&http.Client{})

		tokenValidator, err := idtoken.NewValidator(context.Background(), clientOption)
		if err != nil {
			return err
		}

		googleClientId := os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
		payload, err := tokenValidator.Validate(context.Background(), token, googleClientId)
		if err != nil {
			data := echo.Map{"message": "Invalid Token"}
			return c.JSON(http.StatusUnauthorized, data)
		}

		user := &User{
			Name:     payload.Claims["name"].(string),
			Email:    payload.Claims["email"].(string),
			ImageUrl: payload.Claims["picture"].(string),
		}
		if err = user.Login(); err != nil {
			data := echo.Map{"message": "Failed to login"}
			return c.JSON(http.StatusUnauthorized, data)
		}

		ctx := c.Request().WithContext(context.WithValue(c.Request().Context(), userIdKey, user.ID))
		c.SetRequest(ctx)
		next(c)
		return nil
	}
}
