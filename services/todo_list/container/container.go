package container

import (
	"github.com/sajjad1993/todo/pkg/log"
	"github.com/sajjad1993/todo/pkg/meesage_broker"
	"github.com/sajjad1993/todo/services/todo_list/adapter/broker/command_handlers"
	"github.com/sajjad1993/todo/services/todo_list/adapter/grpc"
	"github.com/sajjad1993/todo/services/todo_list/adapter/reposiroty/orm"
	"github.com/sajjad1993/todo/services/todo_list/app"
	"github.com/sajjad1993/todo/services/todo_list/config"
	"github.com/sajjad1993/todo/services/todo_list/domain/todo"
	"gorm.io/gorm"
)

type Container struct {
	Logger          log.Logger
	Config          config.Config
	TodoService     app.UseCase
	TodoRepository  todo.Repository
	Consumer        meesage_broker.Consumer
	CommandsHandler *command_handlers.CommandsHandler
	Database        *gorm.DB
	Handler         *grpc.Handler
}

func NewContainer(Logger log.Logger, Config config.Config, TodoService app.UseCase, TodoRepository todo.Repository,
	Consumer meesage_broker.Consumer, CommandsHandler *command_handlers.CommandsHandler, Handler *grpc.Handler,
	Database *gorm.DB,
) (*Container, error) {
	orm.Migrate(Database)
	return &Container{
		Logger:          Logger,
		Config:          Config,
		TodoService:     TodoService,
		Consumer:        Consumer,
		TodoRepository:  TodoRepository,
		CommandsHandler: CommandsHandler,
		Handler:         Handler,
	}, nil

}
