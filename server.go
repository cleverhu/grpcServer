package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpcServer/services"
	"log"
	"net/http"
)

func main() {
	fmt.Println("start server")
	creds, err := credentials.NewServerTLSFromFile("keys/server.crt", "keys/server_no_password.key")
	if err != nil {
		log.Fatal(err)
	}


	rpcServer := grpc.NewServer(grpc.Creds(creds))
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Proto)
		fmt.Println(request)
		rpcServer.ServeHTTP(writer, request)
	})

	httpServer := &http.Server{Addr: ":8081", Handler: mux}
	log.Fatal(httpServer.ListenAndServeTLS("keys/server.crt", "keys/server_no_password.key"))
}
