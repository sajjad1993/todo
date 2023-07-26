// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package initilizer

import (
	"context"
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

// Injectors from wire.go:

// InitializeContainer  is dependency injected form of having *service.Container
func InitializeContainer(ctx context.Context) (*container.Container, error) {
	logger := log.NewLogger()
	configConfig, err := config.New()
	if err != nil {
		return nil, err
	}
	dbConfig := config.NewDatabaseConfig(configConfig)
	gormDB, err := db.NewDb(dbConfig)
	if err != nil {
		return nil, err
	}
	repository := orm.NewTodoRepository(gormDB)
	useCase := app.NewService(repository)
	meesage_brokerConfig := config.NewMessageBrokerConfig(configConfig)
	consumer, err := meesage_broker.NewConsumer(meesage_brokerConfig)
	if err != nil {
		return nil, err
	}
	producer, err := meesage_broker.NewProducer(meesage_brokerConfig)
	if err != nil {
		return nil, err
	}
	commandPublisher := publisher.New(producer)
	createTodoListHandler := command_handlers.NewCreateTodoListHandler(consumer, useCase, logger, commandPublisher)
	createTodoHandler := command_handlers.NewCreateTodoHandler(consumer, useCase, logger, commandPublisher)
	deleteTodoListHandler := command_handlers.NewDeleteTodoListHandler(consumer, useCase, logger, commandPublisher)
	updateTodoListHandler := command_handlers.NewUpdateTodoListHandler(consumer, useCase, logger)
	updateTodoHandler := command_handlers.NewUpdateTodoHandler(consumer, useCase, logger)
	deleteTodoHandler := command_handlers.NewDeleteTodoHandler(consumer, useCase, logger)
	commandsHandler, err := command_handlers.New(createTodoListHandler, createTodoHandler, deleteTodoListHandler, updateTodoListHandler, updateTodoHandler, deleteTodoHandler, logger)
	if err != nil {
		return nil, err
	}
	handler := grpc.New(useCase, logger)
	containerContainer, err := container.NewContainer(logger, configConfig, useCase, repository, consumer, commandsHandler, handler, gormDB)
	if err != nil {
		return nil, err
	}
	return containerContainer, nil
}
