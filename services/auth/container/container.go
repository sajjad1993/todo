package container

import (
	"github.com/sajjad1993/todo/pkg/log"
	"github.com/sajjad1993/todo/services/auth/adapter/grpc"
	"github.com/sajjad1993/todo/services/auth/config"
	"github.com/sajjad1993/todo/services/auth/domain/user"
	"github.com/sajjad1993/todo/services/user/app"
)

type Container struct {
	Logger         log.Logger
	Config         config.Config
	UserRepository user.Repository
	Handler        *grpc.Handler
	app            app.UseCase
}

func NewContainer(
	Logger log.Logger,
	Config config.Config,
	UserRepository user.Repository,
	Handler *grpc.Handler,
) (*Container, error) {
	return &Container{
		Logger:         Logger,
		Config:         Config,
		UserRepository: UserRepository,
		Handler:        Handler,
	}, nil

}
