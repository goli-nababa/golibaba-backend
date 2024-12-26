package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/goli-nababa/golibaba-backend/common"
	"github.com/goli-nababa/golibaba-backend/modules/gateway_client"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"user_service/app"
	"user_service/config"

	pb "github.com/goli-nababa/golibaba-backend/proto/pb"
	server "user_service/api/grpc"
	"user_service/api/http"
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

	log.Println("Registering service to gateway")

	gateway := gateway_client.NewGatewayClient(c.Services["gateway"], 1)

	heartBeat := gateway_client.HeartBeat{
		Url: c.Info.HeartBeat.Url,
		TTL: int64(c.Info.HeartBeat.TTL),
	}

	err = gateway.RegisterService(gateway_client.RegisterRequest{
		Name:      c.Info.Name,
		Version:   c.Info.Version,
		UrlPrefix: c.Info.UrlPrefix,
		Host:      c.Server.Host,
		Port:      strconv.Itoa(int(c.Server.Port)),
		BaseUrl:   c.Info.BaseUrl,
		Mapping: map[string]gateway_client.Endpoint{
			"/login": {
				Url: "/account/login",
				PermissionList: map[string]any{
					"super_admin": append(common.RbacAdminPermissions, "user_service:user:delete"),
				},
			},
		},
		HeartBeat: heartBeat,
	})

	if err != nil {
		return
	}

	log.Println("Starting gRPC Server on port 8081")
	err = grpcServer.Serve(l)

	if err != nil {
		return
	}
}
