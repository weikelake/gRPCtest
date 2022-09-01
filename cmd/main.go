package main

import (
	"context"
	"gRPCtest/pkg/proto"
	"gRPCtest/src"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

func main() {
	err := initConfig()
	if err != nil {
		log.Fatalln(err)
	}
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
	log.Printf("rest api server listening at %s", viper.GetString("restPort"))
	if err := http.ListenAndServe(viper.GetString("restPort"), mux); err != nil {
		panic(err)
	}
}

func runGrpc() {
	s := grpc.NewServer()
	srv := &src.Server{}
	api.RegisterRusProfileServer(s, srv)

	log.Printf("grpc api server listening at %s", viper.GetString("grpcPort"))
	l, err := net.Listen("tcp", viper.GetString("grpcPort"))
	if err != nil {
		log.Fatal(err)
	}
	err = s.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
