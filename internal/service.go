package internal

import "github.com/ardafirdausr/todo-server/internal/entity"

type SSOAuthenticator interface {
	Authenticate(token string) (*entity.User, error)
}

type Tokenizer interface {
	Generate(entity.TokenPayload) (string, error)
}
