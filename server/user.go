package main

import (
	"main/database"
	gen_ "main/z_generated"
	"context"
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"
)

type server struct{}

type User struct {
	Name     string `json:"name"`
    Class    string `json:"class"`
    RollNum  int64  `json:"roll_num"`
    Metadata []byte `json:"metadata"`
}

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

	gen_.RegisterAddServiceServer(srv, &server{})

	srv.Serve(listener)

}

func (s *server) SetUser(ctx context.Context, user *gen_.User) (*gen_.User, error) {
	name := user.GetName()
	class := user.GetClass()
	roll_num := user.GetRollNum()
	metadata := user.GetMetadata()

	key := "Paranjaya" + ":" + name + ":" + class + ":" + string(roll_num)

	newUser := database.User{name, class, roll_num, metadata}
	val, err := db.Set(key, newUser)

	return &gen_.User{Name: name, Class: class, RollNum: roll_num, Metadata: metadata}, err
}

func (s *server) GetUser(ctx context.Context, user *gen_.User) (*gen_.User, error) {
	name := user.GetName()
	class := user.GetClass()
	roll_num := user.GetRollNum()

	key := "Paranjaya" + ":" + name + ":" + class + ":" + string(roll_num)

	newUser, err := db.Get(key)
	return &gen_.User{Name: newUser.Name, Class: newUser.Class, RollNum: newUser.RollNum, Metadata: newUser.Metadata}, err
}
