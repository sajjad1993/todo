package command_handlers

import (
	"context"
	"encoding/json"
	"github.com/sajjad1993/todo/internal/todo_list/app"
	"github.com/sajjad1993/todo/internal/todo_list/domain/todo"
	"github.com/sajjad1993/todo/pkg/errs"
	"github.com/sajjad1993/todo/pkg/log"
	"github.com/sajjad1993/todo/pkg/meesage_broker"
	"time"
)

const CreateTodoList = "CREATE_TODO_LIST"

type CreateTodoListHandler struct {
	timeOut  time.Duration
	key      string
	consumer meesage_broker.Consumer
	service  app.UseCase
	logger   log.Logger
}

func (h CreateTodoListHandler) Handle() error {
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

func (h *CreateTodoListHandler) handleService(data []byte) error {
	var ent todo.List
	err := json.Unmarshal(data, &ent)
	if err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), h.timeOut)

	err = h.service.CreateToDoList(ctx, &ent)
	if err != nil {
		return err
	}
	return nil
}

func NewCreateTodoListHandler(consumer meesage_broker.Consumer, service app.UseCase, logger log.Logger) *CreateTodoListHandler {
	timeout := 5 * time.Second //todo move to config
	key := CreateTodoList
	return &CreateTodoListHandler{
		timeOut:  timeout,
		key:      key,
		consumer: consumer,
		service:  service,
		logger:   logger,
	}
}
