package broker

import (
	"context"
	"encoding/json"
	"github.com/sajjad1993/todo/pkg/errs"
	"github.com/sajjad1993/todo/pkg/meesage_broker"
)

type CommandHandler interface {
	Handle(ctx context.Context, body interface{}, key string) error
}

type RabbitCommandHandler struct {
	producer meesage_broker.Producer
}

func (c *RabbitCommandHandler) Handle(ctx context.Context, body interface{}, key string) error {
	data, err := json.Marshal(body)
	if err != nil {
		return errs.NewInternalError(err.Error())
	}
	return c.producer.Publish(key, data)
}

func New(producer meesage_broker.Producer) CommandHandler {
	return &RabbitCommandHandler{producer: producer}
}
