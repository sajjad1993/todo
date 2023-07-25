package grpc

import (
	"context"
	"github.com/sajjad1993/todo/internal/auth/app"
	"github.com/sajjad1993/todo/internal/auth/app/jwt"
	"github.com/sajjad1993/todo/internal/auth/domain/user"
	"github.com/sajjad1993/todo/internal/common"
	rpc "github.com/sajjad1993/todo/internal/common/rpc/auth/api/protobuf"
	"github.com/sajjad1993/todo/pkg/log"
	"google.golang.org/grpc/status"
	"strconv"
)

type Handler struct {
	rpc.UnimplementedAuthServiceServer
	authService app.UseCase
	jwt         jwt.JWT
	logger      log.Logger
}

func New(authService app.UseCase, logger log.Logger, jwt jwt.JWT) *Handler {
	return &Handler{
		authService: authService,
		logger:      logger,
		jwt:         jwt,
	}
}

func (h *Handler) SignIn(ctx context.Context, request *rpc.SignInRequest) (*rpc.SignInResponse, error) {
	entity := &user.User{
		Email:    request.Email,
		Password: request.Password,
	}
	token, err := h.authService.SignIn(ctx, entity)
	if err != nil {
		return nil, status.Errorf(common.GetGrpcStatusCodeByError(err), err.Error())
	}
	return &rpc.SignInResponse{
		AccessToken: token,
	}, nil
}

func (h *Handler) CheckToken(ctx context.Context, request *rpc.CheckTokenRequest) (*rpc.CheckTokenResponse, error) {

	userToken, err := h.jwt.ParseToken(request.AccessToken)
	if err != nil {
		return nil, err
	}
	id, err := strconv.Atoi(userToken.ID)
	if err != nil {
		return nil, err
	}
	return &rpc.CheckTokenResponse{
		Token: &rpc.Token{
			Name:  userToken.Name,
			Email: userToken.Email,
			ID:    uint64(id),
		},
	}, nil
}
