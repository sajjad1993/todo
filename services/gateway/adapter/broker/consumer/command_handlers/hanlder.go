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
	deleteTodo *command.DeleteTodo,
	updateTodoList *command.UpdateTodoList,
	updateTodo *command.UpdateTodo,

) (*CommandsHandlers, error) {
	commandHandlers := &CommandsHandlers{
		logger: logger,
	}

	signUpHandler := NewCommandHandler(consumer, signUp, logger)
	err := signUpHandler.Handle()
	if err != nil {
		return nil, err
	}
	createTodoListHandler := NewCommandHandler(consumer, createTodoList, logger)
	err = createTodoListHandler.Handle()
	if err != nil {
		return nil, err
	}

	createTodoHandler := NewCommandHandler(consumer, createTodo, logger)
	err = createTodoHandler.Handle()
	if err != nil {
		return nil, err
	}

	deleteTodoListHandler := NewCommandHandler(consumer, deleteTodoList, logger)
	err = deleteTodoListHandler.Handle()
	if err != nil {
		return nil, err
	}

	deleteTodoHandler := NewCommandHandler(consumer, deleteTodo, logger)
	err = deleteTodoHandler.Handle()
	if err != nil {
		return nil, err
	}
	updateTodoListHandler := NewCommandHandler(consumer, updateTodoList, logger)
	err = updateTodoListHandler.Handle()
	if err != nil {
		return nil, err
	}
	updateTodoHandler := NewCommandHandler(consumer, updateTodo, logger)
	err = updateTodoHandler.Handle()
	if err != nil {
		return nil, err
	}
	commandHandlers.handlers = append(commandHandlers.handlers,
		signUpHandler,
		createTodoHandler,
		deleteTodoListHandler,
		deleteTodoHandler,
		signUpHandler,
		updateTodoListHandler,
		updateTodoHandler,
	)

	return commandHandlers, nil
}
