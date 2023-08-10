// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package initilizer

import (
	"context"
	"github.com/sajjad1993/todo/pkg/log"
	"github.com/sajjad1993/todo/pkg/meesage_broker"
	"github.com/sajjad1993/todo/pkg/meesage_broker/publisher"
	"github.com/sajjad1993/todo/services/gateway/adapter/auth_client"
	"github.com/sajjad1993/todo/services/gateway/adapter/channel_manager"
	"github.com/sajjad1993/todo/services/gateway/adapter/consumer"
	"github.com/sajjad1993/todo/services/gateway/adapter/controller/commands"
	"github.com/sajjad1993/todo/services/gateway/adapter/producer"
	"github.com/sajjad1993/todo/services/gateway/adapter/todo_list_client"
	"github.com/sajjad1993/todo/services/gateway/app"
	"github.com/sajjad1993/todo/services/gateway/app/command"
	"github.com/sajjad1993/todo/services/gateway/app/query"
	"github.com/sajjad1993/todo/services/gateway/config"
	"github.com/sajjad1993/todo/services/gateway/container"
	"github.com/sajjad1993/todo/services/gateway/infrastructure/restapi/handlers"
)

// Injectors from wire.go:

// InitializeContainer  is dependency injected form of having *service.Container
func InitializeContainer(ctx context.Context) (*container.Container, error) {
	logger := log.NewLogger()
	configConfig, err := config.New()
	if err != nil {
		return nil, err
	}
	meesage_brokerConfig := config.NewMessageBrokerConfig(configConfig)
	meesage_brokerProducer, err := meesage_broker.NewProducer(meesage_brokerConfig)
	if err != nil {
		return nil, err
	}
	commandPublisher := publisher.New(meesage_brokerProducer)
	writer := producer.NewUserProducer(commandPublisher)
	signUpHandler := command.NewSignUpCommand(writer)
	todoWriter := producer.NewTodoProducer(commandPublisher)
	createTodoHandler := command.NewCreateTodoCommand(todoWriter)
	createTodoListHandler := command.NewCreateTodoListCommand(todoWriter)
	updateTodoListHandler := command.NewUpdateTodoListCommand(todoWriter)
	deleteTodoListHandler := command.NewDeleteTodoListCommand(todoWriter)
	updateTodoItemHandler := command.NewUpdateTodoItemCommand(todoWriter)
	deleteTodoItemHandler := command.NewDeleteTodoItemCommand(todoWriter)
	appCommands := app.NewCommands(signUpHandler, createTodoHandler, createTodoListHandler, updateTodoListHandler, deleteTodoListHandler, updateTodoItemHandler, deleteTodoItemHandler)
	repository, err := auth_client.New(logger, configConfig)
	if err != nil {
		return nil, err
	}
	signIn := query.NewSignInQuery(repository)
	checkToken := query.NewCheckTokenQuery(repository)
	reader, err := todo_list_client.New(logger, configConfig)
	if err != nil {
		return nil, err
	}
	listToDoList := query.NewListToDoList(reader)
	queries := app.NewQueries(signIn, checkToken, listToDoList)
	application := app.New(appCommands, queries)
	channelCommandManager := channel_manager.NewCommandChannelManager()
	commandsCommands := commands.NewCommandController(appCommands, channelCommandManager)
	handler := handlers.NewHandler(application, commandsCommands)
	meesage_brokerConsumer, err := meesage_broker.NewConsumer(meesage_brokerConfig)
	if err != nil {
		return nil, err
	}
	commandsHandlers, err := consumer.New(logger, meesage_brokerConsumer, channelCommandManager)
	if err != nil {
		return nil, err
	}
	containerContainer, err := container.NewContainer(logger, configConfig, application, commandPublisher, meesage_brokerProducer, handler, commandsHandlers)
	if err != nil {
		return nil, err
	}
	return containerContainer, nil
}
