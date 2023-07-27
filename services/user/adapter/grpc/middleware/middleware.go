package middleware

import (
	"context"
	"google.golang.org/grpc"
)

func GRPCInterceptor() grpc.ServerOption {
	grpcServerOptions := grpc.UnaryInterceptor(func(
		ctx context.Context,
		request interface{},
		serverInfo *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		return handler(ctx, request)
	})

	return grpcServerOptions
}
