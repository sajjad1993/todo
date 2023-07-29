package command_handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/sajjad1993/todo/pkg/errs"
	"github.com/sajjad1993/todo/pkg/log"
	"github.com/sajjad1993/todo/pkg/meesage_broker"
	"github.com/sajjad1993/todo/pkg/meesage_broker/command_utils"
	"github.com/sajjad1993/todo/pkg/meesage_broker/publisher"
	"time"
)

type QueueManager struct {
	timeOut   time.Duration
	key       string
	DoneKey   string
	consumer  meesage_broker.Consumer
	logger    log.Logger
	publisher publisher.CommandPublisher
}

func (h QueueManager) Consume(commandHandler CommandHandler) error {
	err := h.consumer.QueueDeclare(h.key)
	if err != nil {
		return err
	}
	messages, err := h.consumer.Consume(h.key)
	h.logger.Infof("start listening to queue : %s", h.key)

	if err != nil {
		return errs.NewInternalError(err.Error())
	}
	go func() {
		for message := range messages {
			h.logger.Infof("new message : %s", message.Body)
			go func(data []byte) {
				err = h.Process(data, commandHandler)
				if err != nil {
					h.logger.Error(err)
				}

			}(message.Body)
		}
	}()
	return nil
}

func (h *QueueManager) Process(data []byte, commandHandler CommandHandler) error {
	fmt.Printf("new message has reccived from %s queue by todo service \n ", h.key)
	ctx, _ := context.WithTimeout(context.Background(), h.timeOut)
	message, commandError := commandHandler.Handle(ctx, data)
	return h.publish(message, commandError)
}

func (h *QueueManager) publish(message *command_utils.CommandMessage, CommandError error) error {
	message.Status = command_utils.GetCommandStatusFromError(CommandError)
	if CommandError != nil {
		message.Message = CommandError.Error()
	}
	ctx, _ := context.WithTimeout(context.Background(), h.timeOut)
	err := h.publisher.Publish(ctx, message, h.DoneKey)
	if err != nil {
		return err
	}
	fmt.Printf("new message has sent from todo service into %s queue \n --- the messsage is %v  ", h.DoneKey, message)
	return nil
}

func serialize[T any](data []byte) (*T, *command_utils.CommandMessage, error) {
	var ent T
	var message command_utils.CommandMessage
	err := json.Unmarshal(data, &message)
	if err != nil {
		return &ent, &message, err
	}
	err = mapstructure.Decode(message.Data, &ent)
	if err != nil {
		return &ent, &message, err
	}
	return &ent, &message, err
}
func NewQueueManger(consumer meesage_broker.Consumer, logger log.Logger, publisher publisher.CommandPublisher, key, doneKey string) *QueueManager {
	timeout := 5 * time.Second //todo move to config

	return &QueueManager{
		timeOut:   timeout,
		key:       key,
		DoneKey:   doneKey,
		consumer:  consumer,
		logger:    logger,
		publisher: publisher,
	}
}
