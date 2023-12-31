// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package initilizer

import (
	"context"
	"github.com/sajjad1993/todo/pkg/log"
	"github.com/sajjad1993/todo/services/auth/adapter/grpc"
	"github.com/sajjad1993/todo/services/auth/adapter/user_repository"
	"github.com/sajjad1993/todo/services/auth/app"
	"github.com/sajjad1993/todo/services/auth/app/jwt"
	"github.com/sajjad1993/todo/services/auth/config"
	"github.com/sajjad1993/todo/services/auth/container"
)

// Injectors from wire.go:

// InitializeContainer  is dependency injected form of having *service.Container
func InitializeContainer(ctx context.Context) (*container.Container, error) {
	logger := log.NewLogger()
	configConfig, err := config.New()
	if err != nil {
		return nil, err
	}
	repository, err := user_repository.New(logger, configConfig)
	if err != nil {
		return nil, err
	}
	jwtJWT := jwt.NewJWT(configConfig)
	useCase := app.NewService(repository, jwtJWT)
	handler := grpc.New(useCase, logger, jwtJWT)
	containerContainer, err := container.NewContainer(logger, configConfig, repository, handler)
	if err != nil {
		return nil, err
	}
	return containerContainer, nil
}
