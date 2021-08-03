package main

import (
	"cache-service/database"
	gen "cache-service/proto"
	"context"
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"
)

type server struct{}

var db database.Database

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8000")

	if err != nil {
		panic(err)
	}

	fmt.Println("Server is listening at port 8000")

	srv := grpc.NewServer()

	databaseImplementation := os.Args[1]
	db, err = database.Factory(databaseImplementation)
	if err != nil {
		panic(err)
	}

	gen.RegisterAddServiceServer(srv, &server{})

	srv.Serve(listener)

}

func (s *server) Set(ctx context.Context, request *gen.Request) (*gen.Response, error) {
	key := "Paranjaya:" + request.GetKey()
	value, err := db.Set(key, request.GetValue())
	return &gen.Response{Value: value}, err
}

func (s *server) Get(ctx context.Context, request *gen.Request) (*gen.Response, error) {
	key := "Paranjaya:" + request.GetKey()
	value, err := db.Get(key)
	return &gen.Response{Value: value}, err
}
