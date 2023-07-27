package grpc

import (
	"context"
	"github.com/sajjad1993/todo/pkg/log"
	rpc2 "github.com/sajjad1993/todo/pkg/rpc"
	rpc "github.com/sajjad1993/todo/pkg/rpc/proto/user/api/protobuf"
	"github.com/sajjad1993/todo/services/user/app"
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
		return nil, status.Errorf(rpc2.GetGrpcStatusCodeByError(err), err.Error())
	}
	return &rpc.GetUserResponse{
		User: &rpc.User{
			Name:     entity.Name,
			Email:    entity.Email,
			Password: entity.HashedPassword,
			ID:       uint64(entity.ID),
		},
	}, nil
}
