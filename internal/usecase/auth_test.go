package usecase

import (
	"errors"
	"testing"

	"github.com/ardafirdausr/todo-server/internal/entity"
	"github.com/ardafirdausr/todo-server/internal/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestSSO_InvalidToken(t *testing.T) {
	authenticator := new(mocks.SSOAuthenticator)
	authenticator.On("Authenticate", "token").Return(nil, errors.New("Invalid Token"))

	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("GetUserByEmail", "userEmail").Return(nil, errors.New("Failed fetching user"))

	authUsecase := NewAuthUsecase(mockUserRepo)
	user, err := authUsecase.SSO("token", authenticator)
	assert.Nil(t, user)
	assert.NotNil(t, err)
}

func TestSSO_FailedGetUser(t *testing.T) {
	expectedUser := entity.User{
		ID:       primitive.NewObjectID(),
		Name:     "John Doe",
		Email:    "JohnDoe@mail.com",
		ImageUrl: "https://image.com/JohnDoe.png",
	}

	authenticator := new(mocks.SSOAuthenticator)
	authenticator.On("Authenticate", "token").Return(&expectedUser, nil)

	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("GetUserByEmail", expectedUser.Email).Return(nil, errors.New("Failed fetching user"))

	authUsecase := NewAuthUsecase(mockUserRepo)
	user, err := authUsecase.SSO("token", authenticator)
	assert.Nil(t, user)
	assert.NotNil(t, err)
}

func TestSSO_SuccessRegisteringUser(t *testing.T) {
	expectedUser := &entity.User{
		ID:       primitive.NewObjectID(),
		Name:     "John Doe",
		Email:    "JohnDoe@mail.com",
		ImageUrl: "https://image.com/JohnDoe.png",
	}
	createUserParam := entity.CreateUserParam{
		Name:     "John Doe",
		Email:    "JohnDoe@mail.com",
		ImageUrl: "https://image.com/JohnDoe.png",
	}

	authenticator := new(mocks.SSOAuthenticator)
	authenticator.On("Authenticate", "token").Return(expectedUser, nil)

	errNotFound := entity.NewErrNotFound("User not found", nil)
	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("GetUserByEmail", expectedUser.Email).Return(nil, errNotFound)
	mockUserRepo.On("Create", createUserParam).Return(expectedUser, nil)

	authUsecase := NewAuthUsecase(mockUserRepo)
	user, err := authUsecase.SSO("token", authenticator)
	assert.Nil(t, err)
	assert.ObjectsAreEqualValues(expectedUser, user)
}

func TestSSO_FailedRegisteringUser(t *testing.T) {
	expectedUser := &entity.User{
		ID:       primitive.NewObjectID(),
		Name:     "John Doe",
		Email:    "JohnDoe@mail.com",
		ImageUrl: "https://image.com/JohnDoe.png",
	}
	createUserParam := entity.CreateUserParam{
		Name:     "John Doe",
		Email:    "JohnDoe@mail.com",
		ImageUrl: "https://image.com/JohnDoe.png",
	}

	authenticator := new(mocks.SSOAuthenticator)
	authenticator.On("Authenticate", "token").Return(expectedUser, nil)

	errNotFound := entity.NewErrNotFound("User not found", nil)
	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("GetUserByEmail", expectedUser.Email).Return(nil, errNotFound)
	mockUserRepo.On("Create", createUserParam).Return(nil, errors.New("failed create user"))

	authUsecase := NewAuthUsecase(mockUserRepo)
	user, err := authUsecase.SSO("token", authenticator)
	assert.NotNil(t, err)
	assert.Nil(t, user)
}

func TestSSO_SuccessLogin(t *testing.T) {
	expectedUser := &entity.User{
		ID:       primitive.NewObjectID(),
		Name:     "John Doe",
		Email:    "JohnDoe@mail.com",
		ImageUrl: "https://image.com/JohnDoe.png",
	}

	authenticator := new(mocks.SSOAuthenticator)
	authenticator.On("Authenticate", "token").Return(expectedUser, nil)

	mockUserRepo := new(mocks.UserRepository)
	mockUserRepo.On("GetUserByEmail", expectedUser.Email).Return(expectedUser, nil)

	authUsecase := NewAuthUsecase(mockUserRepo)
	user, err := authUsecase.SSO("token", authenticator)
	assert.Nil(t, err)
	assert.ObjectsAreEqualValues(expectedUser, user)
}

func TestGenerateAuthToken_Failed(t *testing.T) {
	user := entity.User{
		ID:       primitive.NewObjectID(),
		Name:     "John Doe",
		Email:    "JohnDoe@mail.com",
		ImageUrl: "https://image.com/JohnDoe.png",
	}
	tokenPayload := entity.TokenPayload{}
	tokenPayload.ID = user.ID
	tokenPayload.Name = user.Name
	tokenPayload.Email = user.Email
	tokenPayload.Imageurl = user.ImageUrl

	expectedToken := "r4nDom-t0k3n-5tr1ngs"

	authenticator := new(mocks.Tokenizer)
	authenticator.On("Generate", tokenPayload).Return("", errors.New("Failed creating token"))

	mockUserRepo := new(mocks.UserRepository)

	authUsecase := NewAuthUsecase(mockUserRepo)
	token, err := authUsecase.GenerateAuthToken(user, authenticator)
	assert.NotNil(t, err)
	assert.NotEqual(t, expectedToken, token)
}

func TestGenerateAuthToken_Success(t *testing.T) {
	user := entity.User{
		ID:       primitive.NewObjectID(),
		Name:     "John Doe",
		Email:    "JohnDoe@mail.com",
		ImageUrl: "https://image.com/JohnDoe.png",
	}
	tokenPayload := entity.TokenPayload{}
	tokenPayload.ID = user.ID
	tokenPayload.Name = user.Name
	tokenPayload.Email = user.Email
	tokenPayload.Imageurl = user.ImageUrl

	expectedToken := "r4nDom-t0k3n-5tr1ngs"

	authenticator := new(mocks.Tokenizer)
	authenticator.On("Generate", tokenPayload).Return(expectedToken, nil)

	mockUserRepo := new(mocks.UserRepository)

	authUsecase := NewAuthUsecase(mockUserRepo)
	token, err := authUsecase.GenerateAuthToken(user, authenticator)
	assert.Nil(t, err)
	assert.Equal(t, expectedToken, token)
}
