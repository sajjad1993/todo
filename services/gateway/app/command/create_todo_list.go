package command

import (
	"context"
	"github.com/sajjad1993/todo/pkg/meesage_broker/broker_utils"
	"github.com/sajjad1993/todo/services/gateway/domain/todo"
)

type CreateTodoList struct {
	TodoList todo.List
	Name     string
	DoneName string
}

func (c *CreateTodoList) GetName() string {
	return c.Name
}

func (c *CreateTodoList) GetDoneName() string {
	return c.DoneName
}

type CreateTodoListHandler CommandHandler[CreateTodoList]

type createTodoListHandler struct {
	todoWriter todo.Writer
}

func (c *createTodoListHandler) Handle(ctx context.Context, cmd CreateTodoList) error {
	err := c.todoWriter.CreateList(ctx, &cmd.TodoList)
	if err != nil {
		return err
	}
	return nil
}

func NewCreateTodoListCommand(todoWriter todo.Writer) CreateTodoListHandler {

	return &createTodoListHandler{
		todoWriter: todoWriter,
	}
}

func NewCreateTodoList(todoList todo.List) *CreateTodoList {

	return &CreateTodoList{
		Name:     broker_utils.CreateTodoListCommand,
		DoneName: broker_utils.DoneCreateTodoListCommand,
		TodoList: todoList,
	}

}
