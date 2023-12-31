package producer

import (
	"context"
	"github.com/sajjad1993/todo/pkg/meesage_broker/broker_utils"
	"github.com/sajjad1993/todo/pkg/meesage_broker/command_utils"
	"github.com/sajjad1993/todo/pkg/meesage_broker/publisher"
	"github.com/sajjad1993/todo/services/gateway/domain/todo"
)

type TodoWriter struct {
	publisher publisher.CommandPublisher
}

func (u *TodoWriter) CreateList(ctx context.Context, todoList *todo.List) error {
	commandMessage := command_utils.NewCommandMessage("", command_utils.SuccessStatus, todoList)
	return u.produce(ctx, commandMessage)
}

func (u *TodoWriter) CreateItem(ctx context.Context, todo *todo.Item) error {
	commandMessage := command_utils.NewCommandMessage("", command_utils.SuccessStatus, todo)
	return u.produce(ctx, commandMessage)
}

func (u *TodoWriter) DeleteItem(ctx context.Context, itemId, userId uint) error {
	message := broker_utils.DeleteTodoMessage{
		ID:     itemId,
		UserID: userId,
	}
	commandMessage := command_utils.NewCommandMessage("", command_utils.SuccessStatus, message)
	return u.produce(ctx, commandMessage)
}

func (u *TodoWriter) UpdateItem(ctx context.Context, id uint, todo *todo.Item) error {
	commandMessage := command_utils.NewCommandMessage("", command_utils.SuccessStatus, todo)
	return u.produce(ctx, commandMessage)
}

func (u *TodoWriter) UpdateList(ctx context.Context, id uint, todoList *todo.List) error {
	commandMessage := command_utils.NewCommandMessage("", command_utils.SuccessStatus, todoList)
	return u.produce(ctx, commandMessage)
}

func (u *TodoWriter) DeleteList(ctx context.Context, listId uint, userId uint) error {
	message := broker_utils.DeleteTodoListMessage{
		ID:     listId,
		UserID: userId,
	}
	commandMessage := command_utils.NewCommandMessage("", command_utils.SuccessStatus, message)
	return u.produce(ctx, commandMessage)
}

func (u *TodoWriter) produce(ctx context.Context, commandMessage *command_utils.CommandMessage) error {
	hash, err := getKeyString(ctx, command_utils.RequestHashKey)
	if err != nil {
		return err
	}
	commandMessage.UpdateHash(hash)
	name, err := getKeyString(ctx, command_utils.CommandNameKey)
	if err != nil {
		return err
	}
	err = u.publisher.Publish(ctx, commandMessage, name)
	if err != nil {
		return err
	}
	return nil
}

func NewTodoProducer(publisher publisher.CommandPublisher) todo.Writer {
	return &TodoWriter{publisher: publisher}
}
