package command_handlers

import "github.com/sajjad1993/todo/pkg/log"

type CommandHandler interface {
	Handle() error
}

type CommandsHandler struct {
	logger                log.Logger
	CreateTodoListHandler *CreateTodoListHandler
	CreateTodoHandler     *CreateTodoHandler
	DeleteTodoListHandler *DeleteTodoListHandler
	UpdateTodoListHandler *UpdateTodoListHandler
	UpdateTodoHandler     *UpdateTodoHandler
	DeleteTodoHandler     *DeleteTodoHandler
}

func New(createTodoListHandler *CreateTodoListHandler, CreateTodoHandler *CreateTodoHandler,
	DeleteTodoListHandler *DeleteTodoListHandler,
	UpdateTodoListHandler *UpdateTodoListHandler,
	UpdateTodoHandler *UpdateTodoHandler, DeleteTodoHandler *DeleteTodoHandler,
	logger log.Logger) (*CommandsHandler, error) {
	handlers := &CommandsHandler{
		CreateTodoListHandler: createTodoListHandler,
		CreateTodoHandler:     CreateTodoHandler,
		DeleteTodoListHandler: DeleteTodoListHandler,
		UpdateTodoListHandler: UpdateTodoListHandler,
		UpdateTodoHandler:     UpdateTodoHandler,
		DeleteTodoHandler:     DeleteTodoHandler,
		logger:                logger,
	}
	err := handlers.CreateTodoListHandler.Handle()
	if err != nil {
		return nil, err
	}
	err = handlers.CreateTodoHandler.Handle()
	if err != nil {
		return nil, err
	}
	err = handlers.DeleteTodoListHandler.Handle()
	if err != nil {
		return nil, err
	}
	err = handlers.UpdateTodoListHandler.Handle()
	if err != nil {
		return nil, err
	}
	err = handlers.UpdateTodoHandler.Handle()
	if err != nil {
		return nil, err
	}
	err = handlers.DeleteTodoHandler.Handle()
	if err != nil {
		return nil, err
	}
	return handlers, nil
}
