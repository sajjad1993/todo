package command

import (
	"context"
	"github.com/sajjad1993/todo/internal/gateway/adapter/broker"
	"github.com/sajjad1993/todo/internal/gateway/domain/todo"
)

const CreateTodoCommand = "CREATE_TODO"

type CreateTodo struct {
	Name    string
	handler broker.CommandHandler
}

func (c *CreateTodo) GetName() string {
	return c.Name
}

func (c *CreateTodo) Execute(ctx context.Context, todoItem *todo.Item) error {

	err := c.handler.Handle(ctx, todoItem, c.GetName())
	if err != nil {
		//we can retry that .
		return err
	}
	return nil
}

func NewCreateTodoCommand(handler broker.CommandHandler) *CreateTodo {
	return &CreateTodo{
		Name:    CreateTodoCommand,
		handler: handler,
	}
}
