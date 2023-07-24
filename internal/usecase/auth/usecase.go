package auth

import (
	"context"
	"github.com/sajjad1993/todo/internal/domain/user"
)

type UseCase interface {
	// SignUp creates a new user in user repository
	SignUp(ctx context.Context, user user.Entity) (*user.Entity, error)
	// SignIn tries to sign in user with given username and password, and returns an access token on success.
	SignIn(ctx context.Context, user user.Entity) (string, error)
}
