package publisher

import (
	"context"
	"encoding/json"
	"github.com/sajjad1993/todo/pkg/errs"
	"github.com/sajjad1993/todo/pkg/meesage_broker"
)

type CommandPublisher interface {
	Publish(ctx context.Context, body interface{}, key string) error
}

type RabbitCommandPublisher struct {
	producer meesage_broker.Producer
}

func (c *RabbitCommandPublisher) Publish(ctx context.Context, body interface{}, key string) error {
	data, err := json.Marshal(body)
	if err != nil {
		return errs.NewInternalError(err.Error())
	}
	err = c.producer.QueueDeclare(key)
	if err != nil {
		return err
	}
	err = c.producer.Publish(key, data)
	return err
}

func New(producer meesage_broker.Producer) CommandPublisher {
	return &RabbitCommandPublisher{producer: producer}
}
