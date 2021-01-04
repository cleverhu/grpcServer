package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpcServer/services"
	"log"
	"net"
)

func main() {
	fmt.Println("start server")
	rpcServer := grpc.NewServer()
	services.RegisterUserServiceServer(rpcServer, new(services.UserService))

	listen, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatal(err)
	}
	rpcServer.Serve(listen)
}
