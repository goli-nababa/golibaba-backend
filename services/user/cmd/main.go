package main

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"user_service/app"
	"user_service/config"

	pb "github.com/goli-nababa/golibaba-backend/proto/pb"
	server "user_service/api/grpc"
)

var configPath = flag.String("config", "config.json", "Path to service config file")

func main() {
	flag.Parse()

	if v := os.Getenv("CONFIG_PATH"); len(v) > 0 {
		*configPath = v
	}

	c := config.MustReadConfig(*configPath)

	l, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatal(err)
	}

	appContainer := app.MustNewApp(c)

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, server.NewUserServiceGRPCApi(appContainer, c))

	log.Println("Starting gRPC Server on port 8081")

	err = grpcServer.Serve(l)

	if err != nil {
		return
	}
}
