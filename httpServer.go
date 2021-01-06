package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"grpcServer/services"
	"log"
	"net/http"
)

func main() {

	gwMux := runtime.NewServeMux()
	opt := []grpc.DialOption{grpc.WithInsecure()}
	err := services.RegisterUserServiceHandlerFromEndpoint(context.Background(), gwMux, ":8088", opt)
	if err != nil {
		log.Fatal(err)
	}

	httpServer := http.Server{Addr: ":8080", Handler: gwMux}
	httpServer.ListenAndServe()
}
