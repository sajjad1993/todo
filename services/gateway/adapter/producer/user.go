package producer

import (
	"context"
	"github.com/sajjad1993/todo/pkg/meesage_broker/command_utils"
	"github.com/sajjad1993/todo/pkg/meesage_broker/publisher"
	"github.com/sajjad1993/todo/services/gateway/domain/user"
)

type UserWriter struct {
	publisher publisher.CommandPublisher
}

func (u *UserWriter) Create(ctx context.Context, userEnt *user.User) error {
	commandMessage := command_utils.NewCommandMessage("", command_utils.SuccessStatus, userEnt)
	return u.produce(ctx, commandMessage)
}

func (u *UserWriter) produce(ctx context.Context, commandMessage *command_utils.CommandMessage) error {
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

func NewUserProducer(publisher publisher.CommandPublisher) user.Writer {
	return &UserWriter{publisher: publisher}
}
