package main

import (
	gen_ "main/z_generated"
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

	client := gen_.NewAddServiceClient(cc)

	request1 := &gen_.User{Name: "Alice", Class: "IV", RollNum: 15, Metadata: []byte{}}
	request2 := &gen_.User{Name: "Bob", Class: "V", RollNum: 20, Metadata: []byte{}}

	setResponse1, _ := client.SetUser(context.Background(), request1)
	setResponse2, _ := client.SetUser(context.Background(), request2)

	getResponse1, _ := client.GetUser(context.Background(), request1)
	getResponse2, _ := client.GetUser(context.Background(), request2)

	fmt.Println("Set Function Response1", setResponse1)
	fmt.Println("Set Function Response2", setResponse2)
	fmt.Println("Get Function Response1", getResponse1)
	fmt.Println("Get Function Response2", getResponse2)

}
