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
	cc, err := grpc.Dial("localhost:8000", opts)

	if err != nil {
		panic(err)
	}

	defer cc.Close()

	client := gen.NewAddServiceClient(cc)

	request1 := &gen.Request{Key: "Name", Value: []byte("Paranjaya Saxena")}
	request2 := &gen.Request{Key: "Age", Value: []byte("22")}

	setResponse1, _ := client.Set(context.Background(), request1)
	setResponse2, _ := client.Set(context.Background(), request2)

	getResponse1, _ := client.Get(context.Background(), request1)
	getResponse2, _ := client.Get(context.Background(), request2)

	fmt.Println("Set Function Response1", setResponse1)
	fmt.Println("Set Function Response2", setResponse2)
	fmt.Println("Get Function Response1", getResponse1)
	fmt.Println("Get Function Response2", getResponse2)

}
