package controller

import (
	"net/http"
	"time"

	"github.com/ardafirdausr/todo-server/internal/app"
	"github.com/ardafirdausr/todo-server/internal/entity"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	services *app.Services
}

func NewAuthController(services *app.Services) *AuthController {
	return &AuthController{services: services}
}

func (ctrl AuthController) Login(c echo.Context) error {
	googleAuth := entity.GoogleAuth{}
	if err := c.Bind(&googleAuth); err != nil {
		return err
	}

	SSOTokenPayload, err := ctrl.services.AuthService.GetGoogleAuthPayload(googleAuth)
	if err != nil {
		payload := echo.Map{"message": err.Error()}
		return c.JSON(http.StatusBadRequest, payload)
	}

	userEmail := SSOTokenPayload.Claims["email"].(string)
	user, err := ctrl.services.AuthService.GetUserByEmail(userEmail)
	c.Logger().Debug(userEmail, user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if user == nil {
		param := entity.CreateUserParam{
			Email:    userEmail,
			Name:     SSOTokenPayload.Claims["name"].(string),
			ImageUrl: SSOTokenPayload.Claims["picture"].(string),
		}
		user, err = ctrl.services.AuthService.Register(param)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}

	JWTPayload := &entity.JWTPayload{}
	JWTPayload.ID = user.ID
	JWTPayload.Name = user.Name
	JWTPayload.Email = user.Email
	JWTPayload.Imageurl = user.ImageUrl
	JWTPayload.ExpiresAt = time.Now().Add(time.Hour * 3).Unix()
	JWTToken, err := ctrl.services.AuthService.GenerateJWTToken(*JWTPayload)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	payload := echo.Map{
		"message": "Login Successful",
		"data":    user,
		"token":   JWTToken,
	}
	return c.JSON(http.StatusOK, payload)
}
