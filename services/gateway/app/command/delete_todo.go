package command

import (
	"context"
	"github.com/sajjad1993/todo/pkg/meesage_broker/broker_utils"
	"github.com/sajjad1993/todo/services/gateway/domain/todo"
)

type DeleteTodoItem struct {
	ID       uint
	UserID   uint
	Name     string
	DoneName string
}

func (c *DeleteTodoItem) GetName() string {
	return c.Name
}

func (c *DeleteTodoItem) GetDoneName() string {
	return c.DoneName
}

type DeleteTodoItemHandler CommandHandler[DeleteTodoItem]

type deleteTodoItemHandler struct {
	todoWriter todo.Writer
}

func (c *deleteTodoItemHandler) Handle(ctx context.Context, cmd DeleteTodoItem) error {
	err := c.todoWriter.DeleteItem(ctx, cmd.ID, cmd.UserID)
	if err != nil {
		return err
	}
	return nil
}

func NewDeleteTodoItemCommand(todoWriter todo.Writer) DeleteTodoItemHandler {

	return &deleteTodoItemHandler{
		todoWriter: todoWriter,
	}
}

func NewDeleteTodoItem(id, userId uint) *DeleteTodoItem {

	return &DeleteTodoItem{
		Name:     broker_utils.DeleteTodoItemCommand,
		DoneName: broker_utils.DoneDeleteTodoItemCommand,
		ID:       id,
		UserID:   userId,
	}

}
