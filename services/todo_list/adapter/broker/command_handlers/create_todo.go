package command_handlers

import (
	"context"
	"github.com/sajjad1993/todo/pkg/log"
	"github.com/sajjad1993/todo/pkg/meesage_broker"
	"github.com/sajjad1993/todo/pkg/meesage_broker/broker_utils"
	"github.com/sajjad1993/todo/pkg/meesage_broker/command_utils"
	"github.com/sajjad1993/todo/pkg/meesage_broker/publisher"
	"github.com/sajjad1993/todo/services/todo_list/app"
	"github.com/sajjad1993/todo/services/todo_list/domain/todo"
)

type CreateTodoHandler struct {
	service app.UseCase
	queue   *QueueManager
	logger  log.Logger
}

func (h CreateTodoHandler) Handle(ctx context.Context, data []byte) (*command_utils.CommandMessage, error) {
	ent, message, err := serialize[todo.Item](data)
	if err != nil {
		return message, err
	}
	commandError := h.service.CreateToDoItem(ctx, ent)
	return message, commandError
}

func NewCreateTodoHandler(consumer meesage_broker.Consumer, service app.UseCase, logger log.Logger, publisher publisher.CommandPublisher) (*CreateTodoHandler, error) {
	queue := NewQueueManger(consumer, logger, publisher, broker_utils.CreateTodoCommand, broker_utils.DONECreateTodoCommand)

	handler := &CreateTodoHandler{
		queue:   queue,
		service: service,
		logger:  logger,
	}
	err := queue.Consume(handler)
	if err != nil {
		return nil, err
	}
	return handler, nil
}
