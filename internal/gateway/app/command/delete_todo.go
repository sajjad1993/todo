package command

import (
	"context"
	"github.com/sajjad1993/todo/internal/gateway/adapter/broker"
)

const DeleteTodoCommand = "DELETE_TODO"

type deleteTodoMessage struct {
	ID     uint
	UserID uint
}
type DeleteTodo struct {
	Name    string
	handler broker.CommandHandler
}

func (c *DeleteTodo) GetName() string {
	return c.Name
}

func (c *DeleteTodo) Execute(ctx context.Context, todoId uint, userId uint) error {
	message := deleteTodoMessage{
		ID:     todoId,
		UserID: userId,
	}
	err := c.handler.Handle(ctx, message, c.GetName())
	if err != nil {
		//we can retry that .
		return err
	}
	return nil
}

func NewDeleteTodoCommand(handler broker.CommandHandler) *DeleteTodo {
	return &DeleteTodo{
		Name:    DeleteTodoCommand,
		handler: handler,
	}
}
