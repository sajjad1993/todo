package command_handlers

import (
	"github.com/sajjad1993/todo/internal/gateway/app/command"
	"github.com/sajjad1993/todo/pkg/log"
	"github.com/sajjad1993/todo/pkg/meesage_broker"
)

type CommandHandler interface {
	Handle() error
}

type CommandsHandlers struct {
	logger   log.Logger
	handlers []CommandHandler
}

func New(logger log.Logger, signUp *command.SignUp, consumer meesage_broker.Consumer) (*CommandsHandlers, error) {
	commandHandlers := &CommandsHandlers{
		logger: logger,
	}

	signUpHandler := NewCommandHandler(consumer, signUp, logger)
	commandHandlers.handlers = append(commandHandlers.handlers, signUpHandler)
	err := signUpHandler.Handle()
	if err != nil {
		return nil, err
	}

	return commandHandlers, nil
}
