package command_handlers

import (
	"encoding/json"
	"fmt"
	"github.com/sajjad1993/todo/pkg/errs"
	"github.com/sajjad1993/todo/pkg/log"
	"github.com/sajjad1993/todo/pkg/meesage_broker"
	"github.com/sajjad1993/todo/pkg/meesage_broker/command_utils"
	"github.com/sajjad1993/todo/services/gateway/app"
)

type ConsumeCommandHandler struct {
	key      string
	consumer meesage_broker.Consumer
	command  app.Command
	logger   log.Logger
}

func (h ConsumeCommandHandler) Handle() error {
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
				err = h.handleService(data)
				if err != nil {
					h.logger.Error(err)
				}

			}(message.Body)
		}
	}()
	return nil
}

func (h *ConsumeCommandHandler) handleService(data []byte) error {
	var message command_utils.CommandMessage
	err := json.Unmarshal(data, &message)
	if err != nil {
		return err
	}
	fmt.Printf("new message has reccived from %s queue by gateway  and the message is %v \n ", h.key, message)

	h.command.DeleteCommandChannel(&message)
	if err != nil {
		return err
	}
	return nil
}

func NewCommandHandler(consumer meesage_broker.Consumer, command app.Command, logger log.Logger) CommandHandler {

	return &ConsumeCommandHandler{
		key:      command.GetDoneName(),
		consumer: consumer,
		command:  command,
		logger:   logger,
	}
}
