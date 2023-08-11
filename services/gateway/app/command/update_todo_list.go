package command

import (
	"context"
	"github.com/sajjad1993/todo/pkg/meesage_broker/broker_utils"
	"github.com/sajjad1993/todo/services/gateway/domain/todo"
)

type UpdateTodoList struct {
	TodoList todo.List
	Name     string
	DoneName string
}

func (c *UpdateTodoList) GetName() string {
	return c.Name
}

func (c *UpdateTodoList) GetDoneName() string {
	return c.DoneName
}

type UpdateTodoListHandler CommandHandler[UpdateTodoList]

type updateTodoListHandler struct {
	todoWriter todo.Writer
}

func (c *updateTodoListHandler) Handle(ctx context.Context, cmd UpdateTodoList) error {
	err := c.todoWriter.CreateList(ctx, &cmd.TodoList)
	if err != nil {
		return err
	}
	return nil
}

func NewUpdateTodoListCommand(todoWriter todo.Writer) UpdateTodoListHandler {

	return &updateTodoListHandler{
		todoWriter: todoWriter,
	}
}

func NewUpdateTodoList(todoList todo.List) *UpdateTodoList {

	return &UpdateTodoList{
		Name:     broker_utils.UpdateTodoListCommand,
		DoneName: broker_utils.DoneUpdateTodoListCommand,
		TodoList: todoList,
	}

}
