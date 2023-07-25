package grpc

import (
	"context"
	"github.com/sajjad1993/todo/internal/common"
	rpc "github.com/sajjad1993/todo/internal/common/rpc/api/protobuf"
	"github.com/sajjad1993/todo/internal/user/app"
	"github.com/sajjad1993/todo/pkg/log"
	"google.golang.org/grpc/status"
)

type Handler struct {
	rpc.UnimplementedUsersServiceServer
	userService app.UseCase
	logger      log.Logger
}

func New(userService app.UseCase, logger log.Logger) *Handler {
	return &Handler{
		userService: userService,
		logger:      logger,
	}
}

func (h *Handler) GetUser(ctx context.Context, request *rpc.GetUserRequest) (*rpc.GetUserResponse, error) {
	entity, err := h.userService.GetUser(ctx, request.Email)
	if err != nil {
		return nil, status.Errorf(common.GetGrpcStatusCodeByError(err), err.Error())
	}
	return &rpc.GetUserResponse{
		Name:     entity.Name,
		Email:    entity.Email,
		Password: entity.HashedPassword,
	}, nil
}
