//go:build wireinject
// +build wireinject

package initilizer

import (
	"context"
	"github.com/google/wire"
	"github.com/sajjad1993/todo/internal/auth/adapter/grpc"
	"github.com/sajjad1993/todo/internal/auth/adapter/user_repository"
	"github.com/sajjad1993/todo/internal/auth/app"
	"github.com/sajjad1993/todo/internal/auth/app/jwt"
	"github.com/sajjad1993/todo/internal/auth/config"
	"github.com/sajjad1993/todo/internal/auth/container"
	"github.com/sajjad1993/todo/pkg/log"
)

// InitializeContainer  is dependency injected form of having *service.Container
func InitializeContainer(ctx context.Context) (*container.Container, error) {
	wire.Build(
		container.NewContainer,
		config.New,
		log.NewLogger,
		user_repository.New,
		app.NewService,
		grpc.New,
		jwt.NewJWT,
	)
	return new(container.Container), nil
}
