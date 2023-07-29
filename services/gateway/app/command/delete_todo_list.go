package command

import (
	"context"
	"fmt"
	"github.com/sajjad1993/todo/pkg/meesage_broker/broker_utils"
	"github.com/sajjad1993/todo/pkg/meesage_broker/command_utils"
	"github.com/sajjad1993/todo/services/gateway/app/publisher"
)

type DeleteTodoList struct {
	Name      string
	DoneName  string
	publisher publisher.CommandPublisher
	*ChannelCommandManager
}

func (c *DeleteTodoList) GetName() string {
	return c.Name
}
func (c *DeleteTodoList) GetDoneName() string {
	return c.DoneName
}
func (c *DeleteTodoList) Execute(ctx context.Context, message *command_utils.CommandMessage) <-chan *command_utils.CommandMessage {
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

func NewDeleteTodoListCommand(publisher publisher.CommandPublisher) *DeleteTodoList {
	return &DeleteTodoList{
		Name:                  broker_utils.DeleteTodoListCommand,
		DoneName:              broker_utils.DoneDeleteTodoListCommand,
		publisher:             publisher,
		ChannelCommandManager: newCommandChannelManager(),
	}
}