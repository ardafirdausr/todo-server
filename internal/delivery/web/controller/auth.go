package controller

import (
	"net/http"
	"os"

	"github.com/ardafirdausr/todo-server/internal/app"
	"github.com/ardafirdausr/todo-server/internal/entity"
	"github.com/ardafirdausr/todo-server/internal/pkg/auth"
	"github.com/ardafirdausr/todo-server/internal/pkg/token"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	usecases *app.Usecases
}

func NewAuthController(usecases *app.Usecases) *AuthController {
	return &AuthController{usecases: usecases}
}

func (ctrl AuthController) Login(c echo.Context) error {
	googleAuth := entity.GoogleAuth{}
	if err := c.Bind(&googleAuth); err != nil {
		return err
	}

	googleSSOClientID := os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
	googleAuthenticator := auth.NewGoogleSSOAuthenticator(googleSSOClientID)
	user, err := ctrl.usecases.AuthUsecase.SSO(googleAuth.TokenID, googleAuthenticator)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	JWTSecretKey := os.Getenv("JWT_SECRET_KEY")
	JWTToknizer := token.NewJWTTokenizer(JWTSecretKey)
	JWTToken, err := ctrl.usecases.AuthUsecase.GenerateAuthToken(*user, JWTToknizer)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := echo.Map{
		"message": "Login Successful",
		"data":    user,
		"token":   JWTToken,
	}
	return c.JSON(http.StatusOK, response)
}
