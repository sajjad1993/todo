package auth

import (
	"context"
	"github.com/sajjad1993/todo/services/gateway/domain/user"
)

type Repository interface {
	Reader
}

type Reader interface {
	GetToken(ctx context.Context, user *user.User) (*Token, error)
	CheckToken(ctx context.Context, token *Token) (*user.User, error)
}
