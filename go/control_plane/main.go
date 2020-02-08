package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	api "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v2"
	"github.com/envoyproxy/go-control-plane/pkg/cache"
	"github.com/envoyproxy/go-control-plane/pkg/server"
	xds "github.com/envoyproxy/go-control-plane/pkg/server"
)

func main() {
	snapshotCache := cache.NewSnapshotCache(false, cache.IDHash{}, nil)
	server := xds.NewServer(context.Background(), snapshotCache, &callbacks{})
	grpcServer := grpc.NewServer()
	lis, _ := net.Listen("tcp", ":8081")

	discovery.RegisterAggregatedDiscoveryServiceServer(grpcServer, server)
	api.RegisterEndpointDiscoveryServiceServer(grpcServer, server)
	api.RegisterClusterDiscoveryServiceServer(grpcServer, server)
	api.RegisterRouteDiscoveryServiceServer(grpcServer, server)
	api.RegisterListenerDiscoveryServiceServer(grpcServer, server)
	grpcServer.Serve(lis)
}

type callbacks struct{}

func (c *callbacks) OnStreamOpen(context.Context, int64, string) error {
	fmt.Println("OnStreamOpen")
	return nil
}

func (c *callbacks) OnStreamClosed(int64) {
	fmt.Println("OnStreamClosed")
}

func (c *callbacks) OnStreamRequest(int64, *api.DiscoveryRequest) error {
	fmt.Println("OnStreamRequest")
	return nil
}

func (c *callbacks) OnStreamResponse(int64, *api.DiscoveryRequest, *api.DiscoveryResponse) {
	fmt.Println("OnStreamResponse")
}

func (c callbacks) OnFetchRequest(context.Context, *api.DiscoveryRequest) error {
	fmt.Println("OnFetchRequest")
	return nil
}

func (c *callbacks) OnFetchResponse(*api.DiscoveryRequest, *api.DiscoveryResponse) {
	fmt.Println("OnFetchResponse")
}

var _ server.Callbacks = &callbacks{}
