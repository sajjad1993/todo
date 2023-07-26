package command

import (
	"context"
	"github.com/sajjad1993/todo/internal/gateway/adapter/broker"
	"github.com/sajjad1993/todo/internal/gateway/domain/todo"
)

const UpdateTodoListCommand = "UPDATE_TODO_LIST"

type UpdateTodoList struct {
	Name    string
	handler broker.CommandHandler
}

func (c *UpdateTodoList) GetName() string {
	return c.Name
}

func (c *UpdateTodoList) Execute(ctx context.Context, todoList *todo.List) error {

	err := c.handler.Handle(ctx, todoList, c.GetName())
	if err != nil {
		//we can retry that .
		return err
	}
	return nil
}

func NewUpdateTodoListCommand(handler broker.CommandHandler) *UpdateTodoList {
	return &UpdateTodoList{
		Name:    UpdateTodoListCommand,
		handler: handler,
	}
}
