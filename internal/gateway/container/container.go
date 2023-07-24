package container

import (
	"github.com/sajjad1993/todo/internal/gateway/adapter/broker"
	"github.com/sajjad1993/todo/internal/gateway/adapter/restapi/handlers"
	"github.com/sajjad1993/todo/internal/gateway/app"
	"github.com/sajjad1993/todo/internal/gateway/config"
	"github.com/sajjad1993/todo/pkg/log"
	"github.com/sajjad1993/todo/pkg/meesage_broker"
)

type Container struct {
	Logger         log.Logger
	Config         config.Config
	Application    *app.Application
	CommandHandler broker.CommandHandler
	Producer       meesage_broker.Producer
	Handler        *handlers.Handler
}

func NewContainer(Logger log.Logger, Config config.Config, Application *app.Application,
	CommandHandler broker.CommandHandler,
	Producer meesage_broker.Producer, Handler *handlers.Handler) (*Container, error) {
	return &Container{
		Logger:         Logger,
		Config:         Config,
		Application:    Application,
		CommandHandler: CommandHandler,
		Producer:       Producer,
		Handler:        Handler,
	}, nil

}
