package command_handlers

import (
	"github.com/sajjad1993/todo/pkg/log"
	"github.com/sajjad1993/todo/pkg/meesage_broker"
	"github.com/sajjad1993/todo/pkg/meesage_broker/broker_utils"
	"github.com/sajjad1993/todo/services/gateway/adapter/channel_manager"
)

type CommandHandler interface {
	Handle() error
}

type CommandsHandlers struct {
	logger   log.Logger
	handlers []CommandHandler
}

func New(logger log.Logger, consumer meesage_broker.Consumer,
	manager *channel_manager.ChannelCommandManager,

) (*CommandsHandlers, error) {
	commandHandlers := &CommandsHandlers{
		logger: logger,
	}

	signUpHandler := NewCommandHandler(consumer, logger, manager, broker_utils.DoneSignUp)
	err := signUpHandler.Handle()
	if err != nil {
		return nil, err
	}
	createTodoListHandler := NewCommandHandler(consumer, logger, manager, broker_utils.DoneCreateTodoListCommand)
	err = createTodoListHandler.Handle()
	if err != nil {
		return nil, err
	}

	createTodoHandler := NewCommandHandler(consumer, logger, manager, broker_utils.DONECreateTodoCommand)
	err = createTodoHandler.Handle()
	if err != nil {
		return nil, err
	}

	deleteTodoListHandler := NewCommandHandler(consumer, logger, manager, broker_utils.DoneDeleteTodoListCommand)
	err = deleteTodoListHandler.Handle()
	if err != nil {
		return nil, err
	}

	deleteTodoHandler := NewCommandHandler(consumer, logger, manager, broker_utils.DoneDeleteTodoItemCommand)
	err = deleteTodoHandler.Handle()
	if err != nil {
		return nil, err
	}
	updateTodoListHandler := NewCommandHandler(consumer, logger, manager, broker_utils.DoneUpdateTodoListCommand)
	err = updateTodoListHandler.Handle()
	if err != nil {
		return nil, err
	}
	updateTodoHandler := NewCommandHandler(consumer, logger, manager, broker_utils.DoneUpdateTodo)
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
