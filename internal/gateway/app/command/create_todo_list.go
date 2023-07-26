package command

import (
	"context"
	"github.com/sajjad1993/todo/internal/gateway/app/publisher"
	"github.com/sajjad1993/todo/internal/gateway/domain/todo"
)

const CreateTodoListCommand = "CREATE_TODO_LIST"

type CreateTodoList struct {
	Name    string
	handler publisher.CommandPublisher
}

func (c *CreateTodoList) GetName() string {
	return c.Name
}

func (c *CreateTodoList) Execute(ctx context.Context, todoList *todo.List) error {

	err := c.handler.Publish(ctx, todoList, c.GetName())
	if err != nil {
		//we can retry that .
		return err
	}
	return nil
}

func NewCreateTodoListCommand(handler publisher.CommandPublisher) *CreateTodoList {
	return &CreateTodoList{
		Name:    CreateTodoListCommand,
		handler: handler,
	}
}
