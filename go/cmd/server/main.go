package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"grpc_example/src/service"
	"log"
	"net"
)

var (
	port    = flag.Int("port", 10001, "The server port")
	address = flag.String("address", "localhost", "The server address")
)

type demoServer struct {
}

func (s *demoServer) GetInfo(ctx context.Context, in *service.Request) (*service.Info, error) {
	info := &service.Info{
		Id:      "test",
		Message: "Hello world",
	}

	if Id := in.GetId(); Id == "0" {
		return info, nil
	} else {
		info.Message = "Other message"
		return info, nil
	}

}

func newDemoServer() *demoServer {
	return &demoServer{}
}

func main() {
	flag.Parse()

	// Create server
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *address, *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	// Create gRPC server
	log.Printf("Starting GRPC server on port %d", *port)
	grpcServer := grpc.NewServer(opts...)

	//  Register our server
	myServer := newDemoServer()
	service.RegisterDemoServer(grpcServer, myServer)

	// Run
	grpcServer.Serve(lis)
}
