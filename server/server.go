package main

import (
	"Example/proto"
	"context"
	"net"

	"google.golang.org/grpc"
)

func main() {

	grpcServer := grpc.NewServer()
	proto.RegisterExampleServiceServer(grpcServer, &ExampleServiceServer{})

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}

type ExampleServiceServer struct {
	proto.UnimplementedExampleServiceServer
}

func (s *ExampleServiceServer) ExampleMethod(ctx context.Context, in *proto.EmptyRequest) (*proto.Response, error) {
	return &proto.Response{Message: "Hello from server"}, nil
}
