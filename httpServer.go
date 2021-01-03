package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpcServer/services"
	"log"
	"net/http"
)

func main() {
	cred, err := credentials.NewClientTLSFromFile("keys/server.crt", "localhost")
	if err != nil {
		log.Fatal(err)
	}

	gwMux := runtime.NewServeMux()
	opt := []grpc.DialOption{grpc.WithTransportCredentials(cred)}
	err = services.RegisterProdServiceHandlerFromEndpoint(context.Background(), gwMux, "localhost:8081", opt)
	if err != nil {
		log.Fatal(err)
	}

	httpServer := http.Server{Addr: ":8080", Handler: gwMux}
	httpServer.ListenAndServe()
}
