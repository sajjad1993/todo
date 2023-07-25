package user_repository

import (
	"context"
	"fmt"
	"github.com/sajjad1993/todo/internal/auth/config"
	"github.com/sajjad1993/todo/internal/auth/domain/user"
	rpc "github.com/sajjad1993/todo/internal/common/rpc/user/api/protobuf"
	"github.com/sajjad1993/todo/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserService struct {
	client rpc.UsersServiceClient
	logger log.Logger
}

func (c *UserService) GetByEmail(ctx context.Context, email string) (*user.User, error) {
	response, err := c.client.GetUser(ctx, &rpc.GetUserRequest{Email: email})
	if err != nil {
		return nil, err
	}
	userEnt := &user.User{
		Email:          response.User.Email,
		HashedPassword: response.User.Password,
		Name:           response.User.Name,
		ID:             uint(response.User.ID),
	}
	return userEnt, nil
}

func New(logger log.Logger, config config.Config) (user.Repository, error) {
	cc, err := grpc.Dial(fmt.Sprintf("%s", config.GetUserServiceAddress()), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal(fmt.Sprintf("cant connect to user service: %s", err))
		cc.Close()
		return nil, err
	}
	client := rpc.NewUsersServiceClient(cc)
	service := UserService{
		client: client,
		logger: logger,
	}
	return &service, nil

}
