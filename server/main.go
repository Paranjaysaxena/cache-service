package main

import (
	gen "cache-service/proto"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type server struct {
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		panic(err)
	}

	fmt.Println("Server is listening at port 50051")

	srv := grpc.NewServer()
	gen.RegisterAddServiceServer(srv, &server{})

	srv.Serve(listener)

}

func (*server) Set(ctx context.Context, request *gen.Request) (*gen.Response, error) {
	return &gen.Response{Value: request.Value}, nil
}

func (*server) Get(ctx context.Context, request *gen.Request) (*gen.Response, error) {
	return &gen.Response{Value: request.Value}, nil
}
