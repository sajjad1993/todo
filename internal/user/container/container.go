package container

import (
	"github.com/sajjad1993/todo/internal/user/adapter/broker/command_handlers"
	grpc "github.com/sajjad1993/todo/internal/user/adapter/grpc"
	"github.com/sajjad1993/todo/internal/user/adapter/reposiroty/orm"
	"github.com/sajjad1993/todo/internal/user/app"
	"github.com/sajjad1993/todo/internal/user/config"
	"github.com/sajjad1993/todo/internal/user/domain/user"
	"github.com/sajjad1993/todo/pkg/log"
	"github.com/sajjad1993/todo/pkg/meesage_broker"
	"gorm.io/gorm"
)

type Container struct {
	Logger          log.Logger
	Config          config.Config
	UserService     app.UseCase
	UserRepository  user.Repository
	Consumer        meesage_broker.Consumer
	Handler         *grpc.Handler
	CommandsHandler *command_handlers.CommandsHandler
	Database        *gorm.DB
}

func NewContainer(Logger log.Logger, Config config.Config, UserService app.UseCase, UserRepository user.Repository,
	Consumer meesage_broker.Consumer, Handler *grpc.Handler, CommandsHandler *command_handlers.CommandsHandler,
	Database *gorm.DB,
) (*Container, error) {
	orm.Migrate(Database)
	return &Container{
		Logger:          Logger,
		Config:          Config,
		UserService:     UserService,
		Consumer:        Consumer,
		UserRepository:  UserRepository,
		Handler:         Handler,
		CommandsHandler: CommandsHandler,
	}, nil

}
