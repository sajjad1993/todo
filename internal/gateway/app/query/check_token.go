package query

import (
	"context"
	"github.com/sajjad1993/todo/internal/gateway/domain/auth"
	"github.com/sajjad1993/todo/internal/gateway/domain/user"
)

type CheckToken struct {
	authRepo auth.Repository
}

func (q *CheckToken) Run(ctx context.Context, token *auth.Token) (*user.User, error) {
	user, err := q.authRepo.CheckToken(ctx, token)
	if err != nil {
		//todo handel grpc errors in grpc layers
		return nil, err
	}
	return user, nil
}

func NewCheckTokenQuery(authRepo auth.Repository) *CheckToken {
	return &CheckToken{
		authRepo: authRepo,
	}
}
