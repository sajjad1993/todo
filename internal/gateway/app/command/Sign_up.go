package command

import (
	"context"
	"github.com/sajjad1993/todo/internal/gateway/adapter/broker"
	"github.com/sajjad1993/todo/internal/gateway/domain/user"
)

const SignUpCommand = "SIGNUP"

type SignUp struct {
	Name    string
	handler broker.CommandHandler
}

func (c *SignUp) GetName() string {
	return c.Name
}

func (c *SignUp) Execute(ctx context.Context, user *user.User) error {
	err := c.handler.Handle(ctx, user, c.GetName())
	if err != nil {
		//we can retry that .
		return err
	}
	return nil
}

func NewSignUpCommand(handler broker.CommandHandler) *SignUp {
	return &SignUp{
		Name:    SignUpCommand,
		handler: handler,
	}
}
