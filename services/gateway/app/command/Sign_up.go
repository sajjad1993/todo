package command

import (
	"context"
	"github.com/sajjad1993/todo/pkg/meesage_broker/broker_utils"
	"github.com/sajjad1993/todo/services/gateway/app"
	"github.com/sajjad1993/todo/services/gateway/domain/user"
)

type SignUp struct {
	User     user.User
	Name     string
	DoneName string
}

func (c *SignUp) GetName() string {
	return c.Name
}

func (c *SignUp) GetDoneName() string {
	return c.DoneName
}

type SignUpHandler app.CommandHandler[SignUp]

type signUpHandler struct {
	userWriter user.Writer
}

func (c *signUpHandler) Handle(ctx context.Context, cmd SignUp) error {
	err := c.userWriter.Create(ctx, &cmd.User)
	if err != nil {
		return err
	}
	return nil
}

func NewSignUpCommand(userWriter user.Writer) SignUpHandler {

	return &signUpHandler{
		userWriter: userWriter,
	}
}

func NewSignUp(user user.User) *SignUp {

	return &SignUp{
		Name:     broker_utils.SignUp,
		DoneName: broker_utils.DoneSignUp,
		User:     user,
	}

}
