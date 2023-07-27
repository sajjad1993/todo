package command_handlers

import (
	"context"
	"encoding/json"
	"github.com/sajjad1993/todo/pkg/errs"
	"github.com/sajjad1993/todo/pkg/log"
	"github.com/sajjad1993/todo/pkg/meesage_broker"
	"github.com/sajjad1993/todo/pkg/meesage_broker/broker_utils"
	"github.com/sajjad1993/todo/services/todo_list/app"
	"github.com/sajjad1993/todo/services/todo_list/domain/todo"
	"time"
)

type UpdateTodoListHandler struct {
	timeOut  time.Duration
	key      string
	consumer meesage_broker.Consumer
	service  app.UseCase
	logger   log.Logger
}

func (h UpdateTodoListHandler) Handle() error {
	err := h.consumer.QueueDeclare(h.key)
	if err != nil {
		return err
	}
	messages, err := h.consumer.Consume(h.key)
	h.logger.Infof("start listening to queue : %s", h.key)

	if err != nil {
		return errs.NewInternalError(err.Error())
	}
	go func() {
		for message := range messages {
			h.logger.Infof("new message : %s", message.Body)
			go func(data []byte) {
				err = h.handleService(data)
				if err != nil {
					h.logger.Error(err)
				}

			}(message.Body)
		}
	}()
	return nil
}

func (h *UpdateTodoListHandler) handleService(data []byte) error {
	var ent todo.List
	err := json.Unmarshal(data, &ent)
	if err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), h.timeOut)

	err = h.service.UpdateToDoList(ctx, ent.ID, &ent)
	if err != nil {
		return err
	}
	return nil
}

func NewUpdateTodoListHandler(consumer meesage_broker.Consumer, service app.UseCase, logger log.Logger) *UpdateTodoListHandler {
	timeout := 5 * time.Second //todo move to config
	return &UpdateTodoListHandler{
		timeOut:  timeout,
		key:      broker_utils.UpdateTodoListCommand,
		consumer: consumer,
		service:  service,
		logger:   logger,
	}
}
