package command_handlers

import (
	"github.com/sajjad1993/todo/pkg/log"
	"github.com/sajjad1993/todo/pkg/meesage_broker/command_utils"
	"golang.org/x/net/context"
)

type CommandHandler interface {
	Handle(ctx context.Context, data []byte) (*command_utils.CommandMessage, error)
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
	return handlers, nil
}
