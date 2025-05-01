package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"user_service/app"
	"user_service/config"

	"github.com/goli-nababa/golibaba-backend/common"
	"github.com/goli-nababa/golibaba-backend/modules/gateway_client"
	"google.golang.org/grpc"

	server "user_service/api/grpc"
	"user_service/api/http"

	pb "github.com/goli-nababa/golibaba-backend/proto/pb"
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

	pb.RegisterUserServiceServer(grpcServer, server.NewUserServiceGRPCApi(appContainer))

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
			"/register": {
				Url: "/account/register",
				PermissionList: map[string]any{
					"super_admin": append(common.RbacAdminPermissions, "user_service:user:delete"),
				},
			},
			"/history": {
				Url: "/dashboard/history",
				PermissionList: map[string]any{
					"super_admin": append(common.RbacAdminPermissions, "user_service:user:delete"),
				},
			},
			"/notifications": {
				Url: "/dashboard/notifications",
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

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	wg := sync.WaitGroup{}
	wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()

		err = http.Bootstrap(appContainer, c.Server)

		if err != nil {
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		log.Println("Starting gRPC Server on port 8081")
		err = grpcServer.Serve(l)

		if err != nil {
			return
		}
	}()

	<-ctx.Done()
	fmt.Println("Server received shutdown signal, waiting for components to stop...")
	return
}
