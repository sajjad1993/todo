package command_handlers

import "github.com/sajjad1993/todo/pkg/log"

type CommandHandler interface {
	Handle() error
}

type CommandsHandler struct {
	logger        log.Logger
	SignUpHandler *SignUpHandler
}

func New(signUpHandler *SignUpHandler, logger log.Logger) (*CommandsHandler, error) {
	handlers := &CommandsHandler{
		SignUpHandler: signUpHandler,
		logger:        logger,
	}
	err := handlers.SignUpHandler.Handle()
	if err != nil {
		return nil, err
	}
	return handlers, nil
}
