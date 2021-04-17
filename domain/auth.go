package domain

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type AuthUsecase interface {
	GenerateToken(username string) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)     // gia tri jwt, truyen &jwt vo
	SignIn(credentials Credentials) (token string, err error) // named return
	SignUp(user *User) (token string, err error)
}


type AuthMiddleware interface {
	AuthorizeJWT() gin.HandlerFunc
}