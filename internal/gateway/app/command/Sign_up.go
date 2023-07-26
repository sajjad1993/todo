package command

import (
	"context"
	"fmt"
	"github.com/sajjad1993/todo/internal/common/broker_utils"
	"github.com/sajjad1993/todo/internal/common/command_utils"
	"github.com/sajjad1993/todo/internal/gateway/app/publisher"
)

type SignUp struct {
	Name            string
	DoneName        string
	publisher       publisher.CommandPublisher
	commandChannels map[string]chan *command_utils.CommandMessage
}

func (c *SignUp) GetName() string {
	return c.Name
}

func (c *SignUp) GetDoneName() string {
	return c.DoneName
}

func (c *SignUp) Execute(ctx context.Context, message *command_utils.CommandMessage) <-chan *command_utils.CommandMessage {
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

func (c *SignUp) SetCommandChannel(commandMessage *command_utils.CommandMessage) chan *command_utils.CommandMessage {

	ch := make(chan *command_utils.CommandMessage)
	c.commandChannels[commandMessage.Hash] = ch
	return ch
}

func (c *SignUp) DeleteCommandChannelOLD(hash string) chan *command_utils.CommandMessage {
	ch := c.commandChannels[hash]
	delete(c.commandChannels, hash)
	return ch
}
func (c *SignUp) DeleteCommandChannel(commandMessage *command_utils.CommandMessage) {
	ch := c.commandChannels[commandMessage.Hash]
	delete(c.commandChannels, commandMessage.Hash)
	ch <- commandMessage
	close(ch)
}

func NewSignUpCommand(publisher publisher.CommandPublisher) *SignUp {

	mapChannel := make(map[string]chan *command_utils.CommandMessage)
	return &SignUp{
		Name:            broker_utils.SignUp,
		DoneName:        broker_utils.DoneSignUp,
		publisher:       publisher,
		commandChannels: mapChannel,
	}
}
