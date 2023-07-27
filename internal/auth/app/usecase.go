package app

import (
	"context"
	"github.com/sajjad1993/todo/internal/auth/domain/user"
)

type UseCase interface {
	SignIn(ctx context.Context, user *user.User) (string, error)
}
