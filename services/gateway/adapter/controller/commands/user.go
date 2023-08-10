package commands

import (
	"context"
	"github.com/sajjad1993/todo/pkg/meesage_broker/command_utils"
	"github.com/sajjad1993/todo/services/gateway/app/command"
	"github.com/sajjad1993/todo/services/gateway/domain/user"
)

func (c *Commands) SignUp(ctx context.Context, userEnt user.User) error {
	cmd := command.NewSignUp(userEnt)
	commandChannel, ctx := c.setContext(ctx, cmd)
	go func() {
		err := c.app.SignUp.Handle(ctx, *cmd)
		if err != nil {
			errMessage := command_utils.NewCommandMessage("", command_utils.GetCommandStatusFromError(err),
				nil)
			c.manager.DeleteCommandChannel(errMessage)
		}
	}()
	return getCommandResult(ctx, commandChannel)
}
