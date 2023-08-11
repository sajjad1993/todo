package command

import (
	"context"
	"github.com/sajjad1993/todo/pkg/meesage_broker/broker_utils"
	"github.com/sajjad1993/todo/services/gateway/domain/todo"
)

type DeleteTodoList struct {
	ID       uint
	UserID   uint
	Name     string
	DoneName string
}

func (c *DeleteTodoList) GetName() string {
	return c.Name
}

func (c *DeleteTodoList) GetDoneName() string {
	return c.DoneName
}

type DeleteTodoListHandler CommandHandler[DeleteTodoList]

type deleteTodoListHandler struct {
	todoWriter todo.Writer
}

func (c *deleteTodoListHandler) Handle(ctx context.Context, cmd DeleteTodoList) error {
	err := c.todoWriter.DeleteList(ctx, cmd.ID, cmd.UserID)
	if err != nil {
		return err
	}
	return nil
}

func NewDeleteTodoListCommand(todoWriter todo.Writer) DeleteTodoListHandler {

	return &deleteTodoListHandler{
		todoWriter: todoWriter,
	}
}

func NewDeleteTodoList(id, userId uint) *DeleteTodoList {

	return &DeleteTodoList{
		Name:     broker_utils.DeleteTodoListCommand,
		DoneName: broker_utils.DoneDeleteTodoListCommand,
		ID:       id,
		UserID:   userId,
	}

}
