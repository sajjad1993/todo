package command_utils

import (
	"errors"
	"github.com/google/uuid"
	"github.com/sajjad1993/todo/pkg/errs"
)

type CommandStatus string

const (
	SuccessStatus             CommandStatus = "success"
	NotFoundErrorStatus       CommandStatus = "not found"
	ValidationErrorStatus     CommandStatus = "validation error"
	DuplicateErrorStatus      CommandStatus = "duplication error"
	UNAuthorizedStatus        CommandStatus = "unauthorized"
	InternalServerErrorStatus CommandStatus = "internal server error"
)
const RequestHashKey = "request-hash"
const CommandNameKey = "command-name"

func GetCommandStatusFromError(err error) CommandStatus {
	if err == nil {
		return SuccessStatus
	}
	if errors.Is(err, errs.ErrValidation) {
		return ValidationErrorStatus
	}
	if errors.Is(err, errs.ErrUnauthorized) {
		return UNAuthorizedStatus
	}
	if errors.Is(err, errs.ErrNotFoundError) {
		return NotFoundErrorStatus
	}
	if errors.Is(err, errs.ErrNoSuchKey) {
		return ValidationErrorStatus
	}
	if errors.Is(err, errs.ErrDuplicateEntity) {
		return DuplicateErrorStatus
	}
	return InternalServerErrorStatus
}

type CommandMessage struct {
	Data    interface{}
	Message string
	Status  CommandStatus
	Hash    string
}

func NewCommandMessage(message string, status CommandStatus, data interface{}) *CommandMessage {
	c := &CommandMessage{
		Data:    data,
		Message: message,
		Status:  status,
	}
	c.setHash()
	return c
}
func (c *CommandMessage) setHash() {
	c.Hash = uuid.NewString()
}

func (c *CommandMessage) UpdateHash(hash string) {
	c.Hash = hash
}

func (c *CommandMessage) GetError() error {
	switch c.Status {
	case SuccessStatus:
		return nil
	case NotFoundErrorStatus:
		return errs.NewNotFoundError(c.Message)
	case ValidationErrorStatus:
		return errs.NewValidationError(c.Message)
	case DuplicateErrorStatus:
		return errs.NewDuplicateEntity(c.Message)
	case UNAuthorizedStatus:
		return errs.NewUnauthorizedError(c.Message)
	case InternalServerErrorStatus:
		return errs.NewInternalError(c.Message)
	default:
		return nil
	}
}
