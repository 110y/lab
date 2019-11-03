package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	echopb "github.com/110y/lab/proto/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort             = 9000
	statusFailedToListen = 1
	statusFailedToServe  = 2
)

type server struct{}

func (s *server) Echo(ctx context.Context, req *echopb.EchoRequest) (*echopb.EchoResponse, error) {
	return &echopb.EchoResponse{
		Host: "foo",
	}, nil
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gs := grpc.NewServer()
	echopb.RegisterEchoServiceServer(gs, &server{})

	reflection.Register(gs)
	service.RegisterChannelzServiceToServer(gs)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		os.Exit(statusFailedToListen)
	}

	go func() {
		if err := gs.Serve(lis); err != nil {
			os.Exit(statusFailedToServe)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM)

	sig := <-sigChan
	switch sig {
	case syscall.SIGTERM:
		gs.GracefulStop()
	}
}
