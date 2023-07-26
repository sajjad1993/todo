package command

import (
	"context"
	"github.com/sajjad1993/todo/internal/gateway/app/publisher"
	"github.com/sajjad1993/todo/internal/gateway/domain/todo"
)

const UpdateTodoCommand = "UPDATE_TODO"

type UpdateTodo struct {
	Name    string
	handler publisher.CommandPublisher
}

func (c *UpdateTodo) GetName() string {
	return c.Name
}

func (c *UpdateTodo) Execute(ctx context.Context, todoItem *todo.Item) error {

	err := c.handler.Publish(ctx, todoItem, c.GetName())
	if err != nil {
		//we can retry that .
		return err
	}
	return nil
}

func NewUpdateTodoCommand(handler publisher.CommandPublisher) *UpdateTodo {
	return &UpdateTodo{
		Name:    UpdateTodoCommand,
		handler: handler,
	}
}
