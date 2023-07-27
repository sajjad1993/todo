package jwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/sajjad1993/todo/services/auth/domain/token"
)

type customClaims struct {
	token.User
	jwt.StandardClaims
}
