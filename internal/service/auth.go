package service

import (
	"context"
	"net/http"
	"os"

	"github.com/ardafirdausr/todo-server/internal"
	"github.com/ardafirdausr/todo-server/internal/entity"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/api/idtoken"
)

type AuthService struct {
	userRepository internal.UserRepository
}

func NewAuthService(userRepository internal.UserRepository) *AuthService {
	return &AuthService{userRepository: userRepository}
}

func (service AuthService) GetUserByEmail(email string) (*entity.User, error) {
	return service.userRepository.GetUserByEmail(email)
}

func (service AuthService) Register(param entity.CreateUserParam) (*entity.User, error) {
	user, err := service.userRepository.Create(param)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (service AuthService) GetGoogleAuthPayload(googleCredential entity.GoogleAuth) (*idtoken.Payload, error) {
	clientOption := idtoken.WithHTTPClient(&http.Client{})
	tokenValidator, err := idtoken.NewValidator(context.Background(), clientOption)
	if err != nil {
		return nil, err
	}

	googleClientId := os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
	payload, err := tokenValidator.Validate(context.Background(), googleCredential.TokenID, googleClientId)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func (service AuthService) GenerateJWTToken(payload entity.JWTPayload) (string, error) {
	JWTSecretKey := os.Getenv("JWT_SECRET_KEY")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	jwtToken, err := token.SignedString([]byte(JWTSecretKey))
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}
