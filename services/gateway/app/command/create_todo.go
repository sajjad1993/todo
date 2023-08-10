package command

import (
	"context"
	"github.com/sajjad1993/todo/pkg/meesage_broker/broker_utils"
	"github.com/sajjad1993/todo/services/gateway/domain/todo"
)

type CreateTodo struct {
	Todo     todo.Item
	Name     string
	DoneName string
}

func (c *CreateTodo) GetName() string {
	return c.Name
}

func (c *CreateTodo) GetDoneName() string {
	return c.DoneName
}

type CreateTodoHandler CommandHandler[CreateTodo]

type createTodoHandler struct {
	todoWriter todo.Writer
}

func (c *createTodoHandler) Handle(ctx context.Context, cmd CreateTodo) error {
	err := c.todoWriter.CreateItem(ctx, &cmd.Todo)
	if err != nil {
		return err
	}
	return nil
}

func NewCreateTodoCommand(todoWriter todo.Writer) CreateTodoHandler {

	return &createTodoHandler{
		todoWriter: todoWriter,
	}
}

func NewCreateTodo(todo todo.Item) *CreateTodo {

	return &CreateTodo{
		Name:     broker_utils.CreateTodoCommand,
		DoneName: broker_utils.DONECreateTodoCommand,
		Todo:     todo,
	}

}
