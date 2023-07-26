package command

import (
	"context"
	"github.com/sajjad1993/todo/internal/gateway/app/publisher"
	"github.com/sajjad1993/todo/internal/gateway/domain/todo"
)

const CreateTodoCommand = "CREATE_TODO"

type CreateTodo struct {
	Name    string
	handler publisher.CommandPublisher
}

func (c *CreateTodo) GetName() string {
	return c.Name
}

func (c *CreateTodo) Execute(ctx context.Context, todoItem *todo.Item) error {

	err := c.handler.Publish(ctx, todoItem, c.GetName())
	if err != nil {
		//we can retry that .
		return err
	}
	return nil
}

func NewCreateTodoCommand(handler publisher.CommandPublisher) *CreateTodo {
	return &CreateTodo{
		Name:    CreateTodoCommand,
		handler: handler,
	}
}
