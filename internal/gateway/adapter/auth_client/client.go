package auth_client

import (
	"context"
	"fmt"
	rpc "github.com/sajjad1993/todo/internal/common/rpc/auth/api/protobuf"
	"github.com/sajjad1993/todo/internal/gateway/config"
	"github.com/sajjad1993/todo/internal/gateway/domain/auth"
	"github.com/sajjad1993/todo/internal/gateway/domain/user"
	"github.com/sajjad1993/todo/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type authService struct {
	client rpc.AuthServiceClient
	logger log.Logger
}

func (c *authService) GetToken(ctx context.Context, user *user.User) (*auth.Token, error) {
	request := &rpc.SignInRequest{
		Email:    user.Email,
		Password: user.Password,
	}
	response, err := c.client.SignIn(ctx, request)
	if err != nil {
		return nil, err
	}
	token := auth.Token(response.AccessToken)
	return &token, nil
}
func (c *authService) CheckToken(ctx context.Context, token *auth.Token) (*user.User, error) {
	request := &rpc.CheckTokenRequest{
		AccessToken: string(*token),
	}
	response, err := c.client.CheckToken(ctx, request)
	if err != nil {
		return nil, err
	}
	entity := user.User{
		Name:  response.Token.Name,
		Email: response.Token.Email,
		ID:    uint(response.Token.ID),
	}
	return &entity, nil
}

func New(logger log.Logger, config config.Config) (auth.Repository, error) {
	cc, err := grpc.Dial(fmt.Sprintf("%s", config.GetAuthServiceAddress()), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Fatal(fmt.Sprintf("cant connect to user service: %s", err))
		cc.Close()
		return nil, err
	}
	client := rpc.NewAuthServiceClient(cc)
	service := authService{
		client: client,
		logger: logger,
	}
	return &service, nil

}
