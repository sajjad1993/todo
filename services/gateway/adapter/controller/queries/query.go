package queries

import (
	"context"
	"github.com/sajjad1993/todo/services/gateway/app"
	"github.com/sajjad1993/todo/services/gateway/domain/auth"
	"github.com/sajjad1993/todo/services/gateway/domain/todo"
	"github.com/sajjad1993/todo/services/gateway/domain/user"
)

type Queries struct {
	app *app.Queries
}

func (q *Queries) SignIn(ctx context.Context, user *user.User) (*auth.Token, error) {
	return q.app.SignIn.Run(ctx, user)
}
func (q *Queries) CheckToken(ctx context.Context, token *auth.Token) (*user.User, error) {
	return q.app.CheckToken.Run(ctx, token)
}
func (q *Queries) ListToDoList(ctx context.Context, userID uint) ([]*todo.List, error) {
	return q.app.ListToDoList.Run(ctx, userID)
}

func NewQueryController(app *app.Queries) *Queries {
	// todo change it to interface later
	return &Queries{
		app: app,
	}
}
