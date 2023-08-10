package container

import (
	"github.com/sajjad1993/todo/pkg/log"
	"github.com/sajjad1993/todo/pkg/meesage_broker"
	"github.com/sajjad1993/todo/pkg/meesage_broker/publisher"
	"github.com/sajjad1993/todo/services/gateway/adapter/broker/consumer/command_handlers"
	"github.com/sajjad1993/todo/services/gateway/app"
	"github.com/sajjad1993/todo/services/gateway/config"
	"github.com/sajjad1993/todo/services/gateway/infrastructure/restapi/handlers"
)

type Container struct {
	Logger           log.Logger
	Config           config.Config
	Application      *app.Application
	CommandPublisher publisher.CommandPublisher
	Producer         meesage_broker.Producer
	Handler          *handlers.Handler
	CommandsHandlers *command_handlers.CommandsHandlers
}

func NewContainer(Logger log.Logger, Config config.Config, Application *app.Application,
	CommandPublisher publisher.CommandPublisher,
	Producer meesage_broker.Producer, Handler *handlers.Handler, CommandsHandlers *command_handlers.CommandsHandlers) (*Container, error) {
	return &Container{
		Logger:           Logger,
		Config:           Config,
		Application:      Application,
		CommandPublisher: CommandPublisher,
		Producer:         Producer,
		Handler:          Handler,
		CommandsHandlers: CommandsHandlers,
	}, nil

}
