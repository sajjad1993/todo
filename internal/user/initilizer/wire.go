//go:build wireinject
// +build wireinject

package initilizer

import (
	"context"
	"github.com/google/wire"
	"github.com/sajjad1993/todo/internal/user/adapter/broker/command_handlers"
	grpc "github.com/sajjad1993/todo/internal/user/adapter/grpc"
	"github.com/sajjad1993/todo/internal/user/adapter/reposiroty/orm"
	"github.com/sajjad1993/todo/internal/user/app"
	"github.com/sajjad1993/todo/internal/user/config"
	"github.com/sajjad1993/todo/internal/user/container"
	"github.com/sajjad1993/todo/pkg/db"
	"github.com/sajjad1993/todo/pkg/log"
	"github.com/sajjad1993/todo/pkg/meesage_broker"
)

// InitializeContainer  is dependency injected form of having *service.Container
func InitializeContainer(ctx context.Context) (*container.Container, error) {
	wire.Build(
		container.NewContainer,
		config.New,
		config.NewDatabaseConfig,
		config.NewMessageBrokerConfig,
		log.NewLogger,
		command_handlers.NewSignUpCommandHandler,
		command_handlers.New,
		app.NewService,
		orm.NewUserRepository,
		grpc.New,
		db.NewDb,
		meesage_broker.NewConsumer,
	)
	return new(container.Container), nil
}
