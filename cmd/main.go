package main

import (
	"context"
	"gRPCtest/pkg/proto"
	"gRPCtest/src"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

func main() {
	go runRest()
	runGrpc()
}

func runRest() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := api.RegisterRusProfileHandlerFromEndpoint(ctx, mux, "localhost:8080", opts)
	if err != nil {
		panic(err)
	}
	log.Printf("server listening at 8090")
	if err := http.ListenAndServe(":8090", mux); err != nil {
		panic(err)
	}
}

func runGrpc() {
	s := grpc.NewServer()
	srv := &src.Server{}
	api.RegisterRusProfileServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	err = s.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}
