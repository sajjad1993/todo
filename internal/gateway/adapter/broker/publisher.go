package broker

import (
	"context"
	"encoding/json"
	"github.com/sajjad1993/todo/internal/gateway/app/publisher"
	"github.com/sajjad1993/todo/pkg/errs"
	"github.com/sajjad1993/todo/pkg/meesage_broker"
)

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

func New(producer meesage_broker.Producer) publisher.CommandPublisher {
	return &RabbitCommandPublisher{producer: producer}
}
