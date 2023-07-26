//go:build wireinject
// +build wireinject

package initilizer

import (
	"context"
	"github.com/google/wire"
	"github.com/sajjad1993/todo/internal/gateway/adapter/auth_client"
	"github.com/sajjad1993/todo/internal/gateway/adapter/broker"
	"github.com/sajjad1993/todo/internal/gateway/adapter/broker/consumer/command_handlers"
	"github.com/sajjad1993/todo/internal/gateway/adapter/restapi/handlers"
	"github.com/sajjad1993/todo/internal/gateway/adapter/todo_list_client"
	"github.com/sajjad1993/todo/internal/gateway/app"
	"github.com/sajjad1993/todo/internal/gateway/app/command"
	"github.com/sajjad1993/todo/internal/gateway/app/query"
	"github.com/sajjad1993/todo/internal/gateway/config"
	"github.com/sajjad1993/todo/internal/gateway/container"
	"github.com/sajjad1993/todo/pkg/log"
	"github.com/sajjad1993/todo/pkg/meesage_broker"
)

// InitializeContainer  is dependency injected form of having *service.Container
func InitializeContainer(ctx context.Context) (*container.Container, error) {
	wire.Build(
		container.NewContainer,
		broker.New,

		config.New,
		log.NewLogger,
		command.NewSignUpCommand,
		app.New,
		meesage_broker.NewProducer,
		meesage_broker.NewConsumer,
		handlers.NewHandler,
		config.NewMessageBrokerConfig,
		query.NewSignInQuery,
		query.NewCheckTokenQuery,
		query.NewListToDoList,
		app.NewCommands,
		app.NewQueries,
		auth_client.New,
		todo_list_client.New,
		command.NewCreateTodoListCommand,
		command.NewCreateTodoCommand,
		command.NewDeleteTodoListCommand,
		command.NewUpdateTodoListCommand,
		command.NewUpdateTodoCommand,
		command.NewDeleteTodoCommand,
		command_handlers.New,
	)
	return new(container.Container), nil
}
