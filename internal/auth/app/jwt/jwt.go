package jwt

import (
	"github.com/sajjad1993/todo/internal/auth/config"
	"github.com/sajjad1993/todo/internal/auth/domain/token"
)

type JWT Access

type Access interface {
	GenerateToken(user *token.User) (string, error)
	ParseToken(accessToken string) (*token.User, error)
}

func NewJWT(cfg config.Config) JWT {
	access := NewAccess(cfg)
	jwt := access
	return jwt

}
