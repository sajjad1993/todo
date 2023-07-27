package app

import (
	"context"
	"github.com/sajjad1993/todo/services/user/domain/user"
)

type UseCase interface {
	// SignUp creates a new user in user repository
	SignUp(ctx context.Context, user *user.User) error
	GetUser(ctx context.Context, email string) (*user.User, error)
}
