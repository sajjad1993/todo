package query

import (
	"context"
	"github.com/sajjad1993/todo/services/gateway/domain/auth"
	"github.com/sajjad1993/todo/services/gateway/domain/user"
)

type SignIn struct {
	authRepo auth.Repository
}

func (q *SignIn) Run(ctx context.Context, user *user.User) (*auth.Token, error) {
	token, err := q.authRepo.GetToken(ctx, user)
	if err != nil {
		//todo handel grpc errors in grpc layers
		return nil, err
	}
	return token, nil
}

func NewSignInQuery(authRepo auth.Repository) *SignIn {
	return &SignIn{
		authRepo: authRepo,
	}
}
