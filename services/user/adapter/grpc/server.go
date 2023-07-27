package grpc

import (
	"fmt"
	"github.com/sajjad1993/todo/pkg/log"
	rpc "github.com/sajjad1993/todo/pkg/rpc/proto/user/api/protobuf"
	"github.com/sajjad1993/todo/services/user/adapter/grpc/middleware"
	"github.com/sajjad1993/todo/services/user/config"
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
	logger.Infof(fmt.Sprintf("Succeesssed to serve gRPC user server: %s", grpcAddress))

	err = grpcServer.Serve(listener)
	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed to serve gRPC server: %v", err))
		panic(fmt.Sprintf("Failed to serve gRPC server: %v", err))
	}
	logger.Infof(fmt.Sprintf("Succeesssed to serve gRPC server: %s", grpcAddress))
}
