package token

import (
	"time"

	"github.com/ardafirdausr/todo-server/internal/entity"
	"github.com/dgrijalva/jwt-go"
)

type JWTTokenizer struct {
	secretKey string
}

func NewJWTTokenizer(secretKey string) JWTTokenizer {
	return JWTTokenizer{secretKey: secretKey}
}

func (JWTT JWTTokenizer) Generate(payload entity.TokenPayload) (string, error) {
	jwtPayload := entity.JWTPayload{}
	jwtPayload.Name = payload.Name
	jwtPayload.Email = payload.Email
	jwtPayload.Imageurl = payload.Imageurl
	jwtPayload.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtPayload)
	jwtToken, err := token.SignedString([]byte(JWTT.secretKey))
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}
