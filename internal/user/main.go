package main

import (
	"context"
	"github.com/sajjad1993/todo/internal/user/adapter/grpc"
	"github.com/sajjad1993/todo/internal/user/container"
	"github.com/sajjad1993/todo/internal/user/initilizer"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// change email
func main() {
	var cc container.Container
	container := &cc
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	container, err := initilizer.InitializeContainer(ctx)
	if err != nil {
		panic(err)
	}

	grpc.Serve(container.Config, container.Logger, container.Handler)
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	container.Logger.Info("todo server shut down gracefully")
}
