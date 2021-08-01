package main

import (
	gen "cache-service/proto"
	"context"
	"fmt"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Hello Client.. ")

	opts := grpc.WithInsecure()
	cc, err := grpc.Dial("localhost:50051", opts)

	if err != nil {
		panic(err)
	}

	defer cc.Close()

	client := gen.NewAddServiceClient(cc)

	request := &gen.Request{Key: "Name", Value: []byte("Paranjaya")}

	setResponse, _ := client.Set(context.Background(), request)

	getResponse, _ := client.Get(context.Background(), request)

	fmt.Println("Set Function Response", setResponse)
	fmt.Println("Get Function Response", getResponse)

}
