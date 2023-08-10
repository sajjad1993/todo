package command

import (
	"context"
	"fmt"
	"github.com/sajjad1993/todo/pkg/meesage_broker/broker_utils"
	"github.com/sajjad1993/todo/pkg/meesage_broker/command_utils"
	"github.com/sajjad1993/todo/services/gateway/adapter/channel_manager"
	"github.com/sajjad1993/todo/services/gateway/app/publisher"
)

type UpdateTodoList struct {
	Name      string
	DoneName  string
	publisher publisher.CommandPublisher
	*channel_manager.ChannelCommandManager
}

func (c *UpdateTodoList) GetDoneName() string {
	return c.DoneName
}
func (c *UpdateTodoList) GetName() string {
	return c.Name
}
func (c *UpdateTodoList) Execute(ctx context.Context, message *command_utils.CommandMessage) <-chan *command_utils.CommandMessage {
	commandChannel := c.SetCommandChannel(message)
	go func() {
		err := c.publisher.Publish(ctx, message, c.GetName())
		if err != nil {
			//we can retry that .
			errMessage := command_utils.NewCommandMessage("", command_utils.GetCommandStatusFromError(err),
				nil)
			c.DeleteCommandChannel(errMessage)
		}
		fmt.Printf("new message has sent from gateway into %s queue \n ", c.GetName())
	}()
	return commandChannel
}

func NewUpdateTodoListCommand(publisher publisher.CommandPublisher) *UpdateTodoList {
	return &UpdateTodoList{
		Name:                  broker_utils.UpdateTodoListCommand,
		DoneName:              broker_utils.DoneUpdateTodoListCommand,
		publisher:             publisher,
		ChannelCommandManager: channel_manager.NewCommandChannelManager(),
	}
}
