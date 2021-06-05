package usecase

import (
	"log"

	"github.com/ardafirdausr/todo-server/internal"
	"github.com/ardafirdausr/todo-server/internal/entity"
)

type AuthUsecase struct {
	userRepository internal.UserRepository
}

func NewAuthUsecase(userRepository internal.UserRepository) *AuthUsecase {
	return &AuthUsecase{userRepository: userRepository}
}

func (service AuthUsecase) SSO(token string, authenticator internal.SSOAuthenticator) (*entity.User, error) {
	reqUser, err := authenticator.Authenticate(token)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	user, err := service.userRepository.GetUserByEmail(reqUser.Email)
	if _, ok := err.(*entity.ErrNotFound); ok {
		param := entity.CreateUserParam{
			Email:    reqUser.Email,
			Name:     reqUser.Name,
			ImageUrl: reqUser.ImageUrl,
		}
		user, err = service.userRepository.Create(param)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return user, nil
}

func (service AuthUsecase) GenerateAuthToken(user entity.User, tokenizer internal.Tokenizer) (string, error) {
	tokenPayload := entity.TokenPayload{}
	tokenPayload.ID = user.ID
	tokenPayload.Name = user.Name
	tokenPayload.Email = user.Email
	tokenPayload.Imageurl = user.ImageUrl
	token, err := tokenizer.Generate(tokenPayload)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	return token, nil
}
