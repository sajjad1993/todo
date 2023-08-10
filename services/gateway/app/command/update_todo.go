package command

import (
	"context"
	"github.com/sajjad1993/todo/pkg/meesage_broker/broker_utils"
	"github.com/sajjad1993/todo/services/gateway/domain/todo"
)

type UpdateTodoItem struct {
	Todo     todo.Item
	Name     string
	DoneName string
}

func (c *UpdateTodoItem) GetName() string {
	return c.Name
}

func (c *UpdateTodoItem) GetDoneName() string {
	return c.DoneName
}

type UpdateTodoItemHandler CommandHandler[UpdateTodoItem]

type updateTodoItemHandler struct {
	todoWriter todo.Writer
}

func (c *updateTodoItemHandler) Handle(ctx context.Context, cmd UpdateTodoItem) error {
	err := c.todoWriter.CreateItem(ctx, &cmd.Todo)
	if err != nil {
		return err
	}
	return nil
}

func NewUpdateTodoItemCommand(todoWriter todo.Writer) UpdateTodoItemHandler {

	return &updateTodoItemHandler{
		todoWriter: todoWriter,
	}
}

func NewUpdateTodoItem(todo todo.Item) *UpdateTodoItem {

	return &UpdateTodoItem{
		Name:     broker_utils.UpdateTodo,
		DoneName: broker_utils.DoneUpdateTodo,
		Todo:     todo,
	}

}
