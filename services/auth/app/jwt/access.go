package jwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/sajjad1993/todo/services/auth/config"
	"github.com/sajjad1993/todo/services/auth/domain/token"
	"time"
)

type TodoAccess struct {
	expiredAt  time.Duration
	signingKey string
}

func NewAccess(cfg config.Config) Access {
	return &TodoAccess{
		signingKey: cfg.GetAccessSignKey(),
		expiredAt:  cfg.GetAccessJWTExp(),
	}
}
func (j TodoAccess) GenerateToken(user *token.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims{
		User: *user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(j.expiredAt).Unix(),
		},
	})
	tokenString, err := token.SignedString([]byte(j.signingKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j TodoAccess) ParseToken(accessToken string) (*token.User, error) {
	token, err := jwt.Parse(
		accessToken,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, &TokenInValid{message: "unexpected signing method"}
			}
			return []byte(j.signingKey), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, &TokenInValid{message: "cannot get claims from token"}
	}
	user, err := getClaims(claims)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func getClaims(claims jwt.MapClaims) (*token.User, error) {

	name, ok := claims["name"].(string)
	if !ok {
		return nil, &TokenInValid{}
	}
	email, ok := claims["email"].(string)
	if !ok {
		return nil, &TokenInValid{}
	}
	id, ok := claims["ID"].(string)
	if !ok {
		return nil, &TokenInValid{}
	}

	user := &token.User{
		Name:  name,
		Email: email,
		ID:    id,
	}

	return user, nil
}
