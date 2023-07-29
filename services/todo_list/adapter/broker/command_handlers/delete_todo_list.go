package command_handlers

import (
	"context"
	"github.com/sajjad1993/todo/pkg/log"
	"github.com/sajjad1993/todo/pkg/meesage_broker"
	"github.com/sajjad1993/todo/pkg/meesage_broker/broker_utils"
	"github.com/sajjad1993/todo/pkg/meesage_broker/command_utils"
	"github.com/sajjad1993/todo/pkg/meesage_broker/publisher"
	"github.com/sajjad1993/todo/services/todo_list/app"
)

type DeleteTodoListHandler struct {
	service app.UseCase
	queue   *QueueManager
	logger  log.Logger
}

func (h DeleteTodoListHandler) Handle(ctx context.Context, data []byte) (*command_utils.CommandMessage, error) {
	ent, message, err := serialize[broker_utils.DeleteTodoListMessage](data)
	if err != nil {
		return message, err
	}
	commandError := h.service.DeleteToDoList(ctx, ent.ID, ent.UserID)
	return message, commandError
}

func NewDeleteTodoListHandler(consumer meesage_broker.Consumer, service app.UseCase, logger log.Logger, publisher publisher.CommandPublisher) (*DeleteTodoListHandler, error) {
	queue := NewQueueManger(consumer, logger, publisher, broker_utils.DeleteTodoListCommand, broker_utils.DoneDeleteTodoListCommand)

	handler := &DeleteTodoListHandler{
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
