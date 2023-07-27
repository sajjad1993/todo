package command

import (
	"context"
	"github.com/sajjad1993/todo/services/gateway/app/publisher"
)

const DeleteTodoCommand = "DELETE_TODO"

type deleteTodoMessage struct {
	ID     uint
	UserID uint
}
type DeleteTodo struct {
	Name    string
	handler publisher.CommandPublisher
}

func (c *DeleteTodo) GetName() string {
	return c.Name
}

func (c *DeleteTodo) Execute(ctx context.Context, todoId uint, userId uint) error {
	message := deleteTodoMessage{
		ID:     todoId,
		UserID: userId,
	}
	err := c.handler.Publish(ctx, message, c.GetName())
	if err != nil {
		//we can retry that .
		return err
	}
	return nil
}

func NewDeleteTodoCommand(handler publisher.CommandPublisher) *DeleteTodo {
	return &DeleteTodo{
		Name:    DeleteTodoCommand,
		handler: handler,
	}
}
