package command

import (
	"context"
	"github.com/sajjad1993/todo/internal/gateway/adapter/broker"
	"github.com/sajjad1993/todo/internal/gateway/domain/todo"
)

const UpdateTodoCommand = "UPDATE_TODO"

type UpdateTodo struct {
	Name    string
	handler broker.CommandHandler
}

func (c *UpdateTodo) GetName() string {
	return c.Name
}

func (c *UpdateTodo) Execute(ctx context.Context, todoItem *todo.Item) error {

	err := c.handler.Handle(ctx, todoItem, c.GetName())
	if err != nil {
		//we can retry that .
		return err
	}
	return nil
}

func NewUpdateTodoCommand(handler broker.CommandHandler) *UpdateTodo {
	return &UpdateTodo{
		Name:    UpdateTodoCommand,
		handler: handler,
	}
}
