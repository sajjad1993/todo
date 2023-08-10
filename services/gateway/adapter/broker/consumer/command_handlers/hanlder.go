package command_handlers

import (
	"github.com/sajjad1993/todo/pkg/log"
	"github.com/sajjad1993/todo/pkg/meesage_broker"
	"github.com/sajjad1993/todo/services/gateway/adapter/channel_manager"
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
	createTodoList *command.CreateTodoList,
	createTodo *command.CreateTodo,
	deleteTodoList *command.DeleteTodoList,
	deleteTodo *command.DeleteTodo,
	updateTodoList *command.UpdateTodoList,
	updateTodo *command.UpdateTodo,
	manager *channel_manager.ChannelCommandManager,

) (*CommandsHandlers, error) {
	commandHandlers := &CommandsHandlers{
		logger: logger,
	}

	signUpHandler := NewCommandHandler(consumer, createTodoList, logger, manager)
	err := signUpHandler.Handle()
	if err != nil {
		return nil, err
	}
	createTodoListHandler := NewCommandHandler(consumer, createTodoList, logger, manager)
	err = createTodoListHandler.Handle()
	if err != nil {
		return nil, err
	}

	createTodoHandler := NewCommandHandler(consumer, createTodo, logger, manager)
	err = createTodoHandler.Handle()
	if err != nil {
		return nil, err
	}

	deleteTodoListHandler := NewCommandHandler(consumer, deleteTodoList, logger, manager)
	err = deleteTodoListHandler.Handle()
	if err != nil {
		return nil, err
	}

	deleteTodoHandler := NewCommandHandler(consumer, deleteTodo, logger, manager)
	err = deleteTodoHandler.Handle()
	if err != nil {
		return nil, err
	}
	updateTodoListHandler := NewCommandHandler(consumer, updateTodoList, logger, manager)
	err = updateTodoListHandler.Handle()
	if err != nil {
		return nil, err
	}
	updateTodoHandler := NewCommandHandler(consumer, updateTodo, logger, manager)
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
