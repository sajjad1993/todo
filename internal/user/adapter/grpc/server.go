package grpc

import (
	"fmt"
	rpc "github.com/sajjad1993/todo/internal/common/rpc/api/protobuf"
	"github.com/sajjad1993/todo/internal/user/adapter/grpc/middleware"
	"github.com/sajjad1993/todo/internal/user/config"
	"github.com/sajjad1993/todo/pkg/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

// Serve serves grpc server on given port
func Serve(config config.Config, logger log.Logger, handler *Handler) {

	go runGRPCServer(config.GetUserServiceAddress(), middleware.GRPCInterceptor(), handler, logger)
}

func runGRPCServer(
	grpcAddress string,
	option grpc.ServerOption,
	handler *Handler,
	logger log.Logger,
) {
	listener, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to listen on %s: %v", grpcAddress, err))
	}

	grpcServer := grpc.NewServer(option)
	reflection.Register(grpcServer)
	rpc.RegisterUsersServiceServer(grpcServer, handler)
	err = grpcServer.Serve(listener)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to serve gRPC server: %v", err))
		panic(fmt.Sprintf("Failed to serve gRPC server: %v", err))
	}
	logger.Infof(fmt.Sprintf("Succeesssed to serve gRPC server: %s", grpcAddress))
}
