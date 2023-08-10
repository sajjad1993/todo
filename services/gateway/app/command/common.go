package command

import (
	"context"
)

type Command interface {
	GetName() string
	GetDoneName() string
}
type CommandHandler[C any] interface {
	Handle(ctx context.Context, cmd C) error
}
