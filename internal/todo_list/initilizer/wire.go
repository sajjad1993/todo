//go:build wireinject
// +build wireinject

package initilizer

import (
	"context"
	"github.com/google/wire"
	"github.com/sajjad1993/todo/internal/common/publisher"
	"github.com/sajjad1993/todo/internal/todo_list/adapter/broker/command_handlers"
	"github.com/sajjad1993/todo/internal/todo_list/adapter/grpc"
	"github.com/sajjad1993/todo/internal/todo_list/adapter/reposiroty/orm"
	"github.com/sajjad1993/todo/internal/todo_list/app"
	"github.com/sajjad1993/todo/internal/todo_list/config"
	"github.com/sajjad1993/todo/internal/todo_list/container"
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
		command_handlers.NewCreateTodoListHandler,
		command_handlers.NewCreateTodoHandler,
		command_handlers.NewDeleteTodoListHandler,
		command_handlers.NewUpdateTodoListHandler,
		command_handlers.NewUpdateTodoHandler,
		command_handlers.NewDeleteTodoHandler,
		grpc.New,
		command_handlers.New,
		app.NewService,
		orm.NewTodoRepository,
		db.NewDb,
		meesage_broker.NewConsumer,
		meesage_broker.NewProducer,
		publisher.New,
	)
	return new(container.Container), nil
}
