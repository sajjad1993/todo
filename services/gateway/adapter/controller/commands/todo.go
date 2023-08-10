package commands

import (
	"context"
	"github.com/sajjad1993/todo/pkg/meesage_broker/command_utils"
	"github.com/sajjad1993/todo/services/gateway/app/command"
	"github.com/sajjad1993/todo/services/gateway/domain/todo"
)

func (c *Commands) CreateTodoList(ctx context.Context, todoList todo.List) error {
	cmd := command.NewCreateTodoList(todoList)
	commandChannel, ctx := c.setContext(ctx, cmd)
	go func() {
		err := c.app.CreateTodoList.Handle(ctx, *cmd)
		if err != nil {
			errMessage := command_utils.NewCommandMessage("", command_utils.GetCommandStatusFromError(err),
				nil)
			c.manager.DeleteCommandChannel(errMessage)
		}
	}()
	return getCommandResult(ctx, commandChannel)
}