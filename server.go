package main

import (
	"google.golang.org/grpc"
	"grpcServer/services"
	"net"
)

func main() {
	rpcServer := grpc.NewServer()
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))

	lis, _ := net.Listen("tcp", ":8081")
	rpcServer.Serve(lis)
}
