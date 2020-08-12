package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"

	pb "github.com/110y/lab/proto/grpcserver"
)

const (
	grpcPort                         = 9000
	statusFailedToListen             = 1
	statusFailedToServe              = 2
	statusFailedToDialGRPCServer2SVC = 3
	envGRPCServer2Endpoint           = "GRPCSERVER2_ENDPOINT"
)

type grpcServer struct{}

var grpcserver2Client pb.InfoClient

func (s *grpcServer) ServerInfo(ctx context.Context, _ *empty.Empty) (*pb.ServerInfoResponse, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get hostname")
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("metadata not found")
	}

	fmt.Printf("%+v\n", md)

	ctx = metadata.NewOutgoingContext(ctx, md)

	res, err := grpcserver2Client.ServerInfo(ctx, &empty.Empty{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get server info from grpcserver2")
	}

	name := fmt.Sprintf("%s - %s", hostname, res.Name)

	return &pb.ServerInfoResponse{Name: name}, nil
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	conn, err := grpc.DialContext(ctx, os.Getenv(envGRPCServer2Endpoint), grpc.WithInsecure())
	if err != nil {
		os.Exit(statusFailedToDialGRPCServer2SVC)
	}

	grpcserver2Client = pb.NewInfoClient(conn)

	gs := grpc.NewServer()
	pb.RegisterInfoServer(gs, &grpcServer{})

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
