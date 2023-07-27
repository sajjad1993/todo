package command_handlers

import (
	"github.com/sajjad1993/todo/pkg/log"
	"github.com/sajjad1993/todo/pkg/meesage_broker"
	"github.com/sajjad1993/todo/services/gateway/app/command"
)

type CommandHandler interface {
	Handle() error
}

type CommandsHandlers struct {
	logger   log.Logger
	handlers []CommandHandler
}

func New(logger log.Logger, consumer meesage_broker.Consumer,
	signUp *command.SignUp,
	createTodoList *command.CreateTodoList,
	createTodo *command.CreateTodo,
	deleteTodoList *command.DeleteTodoList,

) (*CommandsHandlers, error) {
	commandHandlers := &CommandsHandlers{
		logger: logger,
	}

	signUpHandler := NewCommandHandler(consumer, signUp, logger)
	commandHandlers.handlers = append(commandHandlers.handlers, signUpHandler)
	err := signUpHandler.Handle()
	if err != nil {
		return nil, err
	}
	createTodoListHandler := NewCommandHandler(consumer, createTodoList, logger)
	commandHandlers.handlers = append(commandHandlers.handlers, signUpHandler)
	err = createTodoListHandler.Handle()
	if err != nil {
		return nil, err
	}

	createTodoHandler := NewCommandHandler(consumer, createTodo, logger)
	commandHandlers.handlers = append(commandHandlers.handlers, createTodoHandler)
	err = createTodoHandler.Handle()
	if err != nil {
		return nil, err
	}

	deleteTodoListHandler := NewCommandHandler(consumer, deleteTodoList, logger)
	commandHandlers.handlers = append(commandHandlers.handlers, deleteTodoListHandler)
	err = deleteTodoListHandler.Handle()
	if err != nil {
		return nil, err
	}

	return commandHandlers, nil
}
