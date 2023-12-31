package command_handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/sajjad1993/todo/pkg/errs"
	"github.com/sajjad1993/todo/pkg/log"
	"github.com/sajjad1993/todo/pkg/meesage_broker"
	"github.com/sajjad1993/todo/pkg/meesage_broker/broker_utils"
	"github.com/sajjad1993/todo/pkg/meesage_broker/command_utils"
	"github.com/sajjad1993/todo/pkg/meesage_broker/publisher"
	"github.com/sajjad1993/todo/services/user/app"
	"github.com/sajjad1993/todo/services/user/domain/user"
	"time"
)

type SignUpHandler struct {
	timeOut   time.Duration
	key       string
	DoneKey   string
	consumer  meesage_broker.Consumer
	service   app.UseCase
	logger    log.Logger
	publisher publisher.CommandPublisher
}

func (h SignUpHandler) Handle() error {
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

func (h *SignUpHandler) handleService(data []byte) error {
	fmt.Printf("new message has reccived from %s queue by user service \n ", h.key)
	var message command_utils.CommandMessage
	err := json.Unmarshal(data, &message)
	if err != nil {
		return err
	}
	var ent user.User
	err = mapstructure.Decode(message.Data, &ent)
	if err != nil {
		return errs.NewValidationError("data is corrected")
	}
	ctx, _ := context.WithTimeout(context.Background(), h.timeOut)

	commandError := h.service.SignUp(ctx, &ent)
	return h.publish(&message, commandError)
}

func (h *SignUpHandler) publish(message *command_utils.CommandMessage, CommandError error) error {
	message.Status = command_utils.GetCommandStatusFromError(CommandError)
	if CommandError != nil {
		message.Message = CommandError.Error()
	}
	ctx, _ := context.WithTimeout(context.Background(), h.timeOut)
	err := h.publisher.Publish(ctx, message, h.DoneKey)
	if err != nil {
		return err
	}
	fmt.Printf("new message has sent from user service into %s queue \n --- the messsage is %v  ", h.DoneKey, message)
	return nil
}
func NewSignUpCommandHandler(consumer meesage_broker.Consumer, service app.UseCase, logger log.Logger, publisher publisher.CommandPublisher) *SignUpHandler {
	timeout := 30 * time.Second //todo move to config
	return &SignUpHandler{
		timeOut:   timeout,
		key:       broker_utils.SignUp,
		DoneKey:   broker_utils.DoneSignUp,
		consumer:  consumer,
		service:   service,
		logger:    logger,
		publisher: publisher,
	}
}
