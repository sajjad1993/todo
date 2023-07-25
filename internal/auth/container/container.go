package container

import (
	"github.com/sajjad1993/todo/internal/auth/adapter/grpc"
	"github.com/sajjad1993/todo/internal/auth/config"
	"github.com/sajjad1993/todo/internal/auth/domain/user"
	"github.com/sajjad1993/todo/internal/user/app"
	"github.com/sajjad1993/todo/pkg/log"
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
