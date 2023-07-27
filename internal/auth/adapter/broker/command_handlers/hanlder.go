package command_handlers

import "github.com/sajjad1993/todo/pkg/log"

type CommandHandler interface {
	Handle() error
}

type CommandsHandler struct {
	logger log.Logger
}

func New(logger log.Logger) (*CommandsHandler, error) {
	handlers := &CommandsHandler{
		logger: logger,
	}
	return handlers, nil
}
