package command

import (
	"context"
	"github.com/sajjad1993/todo/pkg/meesage_broker/broker_utils"
	"github.com/sajjad1993/todo/services/gateway/app/publisher"
	"github.com/sajjad1993/todo/services/gateway/domain/todo"
)

type UpdateTodoList struct {
	Name    string
	handler publisher.CommandPublisher
}

func (c *UpdateTodoList) GetName() string {
	return c.Name
}

func (c *UpdateTodoList) Execute(ctx context.Context, todoList *todo.List) error {

	err := c.handler.Publish(ctx, todoList, c.GetName())
	if err != nil {
		//we can retry that .
		return err
	}
	return nil
}

func NewUpdateTodoListCommand(handler publisher.CommandPublisher) *UpdateTodoList {
	return &UpdateTodoList{
		Name:    broker_utils.UpdateTodoListCommand,
		handler: handler,
	}
}
