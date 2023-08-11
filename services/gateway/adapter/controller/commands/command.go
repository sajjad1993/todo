package commands

import (
	"context"
	"github.com/sajjad1993/todo/pkg/errs"
	"github.com/sajjad1993/todo/pkg/meesage_broker/command_utils"
	"github.com/sajjad1993/todo/services/gateway/adapter/channel_manager"
	"github.com/sajjad1993/todo/services/gateway/app"
	"github.com/sajjad1993/todo/services/gateway/app/command"
)

type Commands struct {
	app     *app.Commands
	manager *channel_manager.ChannelCommandManager
}

func (c *Commands) setContext(ctx context.Context, cmd command.Command) (<-chan *command_utils.CommandMessage, context.Context) {
	commandMessage := command_utils.NewCommandMessage("", command_utils.SuccessStatus, nil)
	commandChannel := c.manager.SetCommandChannel(commandMessage)
	ctx = context.WithValue(ctx, command_utils.RequestHashKey, commandMessage.Hash)
	ctx = context.WithValue(ctx, command_utils.CommandNameKey, cmd.GetName())
	return commandChannel, ctx
}

func getCommandResult(ctx context.Context, commandChannel <-chan *command_utils.CommandMessage) error {
	select {
	case <-ctx.Done():
		return errs.NewTimeOut(ctx.Err().Error())
	case message := <-commandChannel:
		err := message.GetError()
		if err != nil {
			return err
		}
		return nil
	}
}

func NewCommandController(commands *app.Commands, manager *channel_manager.ChannelCommandManager) *Commands {
	// todo change it to interface later
	return &Commands{
		app:     commands,
		manager: manager,
	}
}
