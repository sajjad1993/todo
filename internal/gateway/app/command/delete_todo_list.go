package command

import (
	"context"
	"github.com/sajjad1993/todo/internal/gateway/app/publisher"
)

const DeleteTodoListCommand = "DELETE_TODO_LIST"

type deleteTodoListMessage struct {
	ID     uint
	UserID uint
}
type DeleteTodoList struct {
	Name    string
	handler publisher.CommandPublisher
}

func (c *DeleteTodoList) GetName() string {
	return c.Name
}

func (c *DeleteTodoList) Execute(ctx context.Context, todoListId uint, userId uint) error {
	message := deleteTodoListMessage{
		ID:     todoListId,
		UserID: userId,
	}
	err := c.handler.Publish(ctx, message, c.GetName())
	if err != nil {
		//we can retry that .
		return err
	}
	return nil
}

func NewDeleteTodoListCommand(handler publisher.CommandPublisher) *DeleteTodoList {
	return &DeleteTodoList{
		Name:    DeleteTodoListCommand,
		handler: handler,
	}
}
