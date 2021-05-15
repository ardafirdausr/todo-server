package entity

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JWTPayload struct {
	ID       primitive.ObjectID `json:"id"`
	Name     string             `json:"name"`
	Email    string             `json:"email"`
	Imageurl string             `json:"imageUrl"`
	jwt.StandardClaims
}

type GoogleAuth struct {
	TokenID  string `json:"token_id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Imageurl string `json:"imageUrl,omitempty"`
}
