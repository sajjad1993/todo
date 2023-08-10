package command

import (
	"context"
	"github.com/sajjad1993/todo/pkg/meesage_broker/command_utils"
)

type CMD interface {
	//todo convert to command after all ...
	GetName() string
	GetDoneName() string
}

type Command interface {
	GetName() string
	GetDoneName() string
	Execute(ctx context.Context, message *command_utils.CommandMessage) <-chan *command_utils.CommandMessage
	SetCommandChannel(commandMessage *command_utils.CommandMessage) chan *command_utils.CommandMessage
	DeleteCommandChannel(commandMessage *command_utils.CommandMessage)
}
type CommandHandler[C any] interface {
	Handle(ctx context.Context, cmd C) error
}
