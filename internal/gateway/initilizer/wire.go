//go:build wireinject
// +build wireinject

package initilizer

import (
	"context"
	"github.com/google/wire"
	"github.com/sajjad1993/todo/internal/gateway/adapter/broker"
	"github.com/sajjad1993/todo/internal/gateway/adapter/restapi/handlers"
	"github.com/sajjad1993/todo/internal/gateway/app"
	"github.com/sajjad1993/todo/internal/gateway/app/command"
	"github.com/sajjad1993/todo/internal/gateway/config"
	"github.com/sajjad1993/todo/internal/gateway/container"
	"github.com/sajjad1993/todo/pkg/log"
	"github.com/sajjad1993/todo/pkg/meesage_broker"
)

// InitializeContainer  is dependency injected form of having *service.Container
func InitializeContainer(ctx context.Context) (*container.Container, error) {
	wire.Build(
		container.NewContainer,
		config.New,
		log.NewLogger,
		command.NewSignUpCommand,
		app.New,
		meesage_broker.New,
		broker.New,
		handlers.NewHandler,
	)
	return new(container.Container), nil
}
