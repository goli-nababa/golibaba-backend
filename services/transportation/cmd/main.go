package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
	grpcApi "transportation/api/grpc"
	"transportation/api/http"

	"transportation/app"
	"transportation/config"
	"transportation/internal/vehicle_request_poller/domain"

	"github.com/goli-nababa/golibaba-backend/modules/gateway_client"
	pb "github.com/goli-nababa/golibaba-backend/proto/pb"
	"google.golang.org/grpc"
)

func main() {
	path := os.Getenv("CONFIG_FILE")
	if path == "" {
		path = "../config.json"
	}

	cfg := config.MustReadConfig(path)

	appContainer := app.NewMustApp(cfg)
	runVehicleRequestsScheduler(appContainer)

	l, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterTripServiceServer(grpcServer, grpcApi.NewTripServiceGRPCApi(appContainer, cfg))

	log.Println("Registering service to gateway")

	gateway := gateway_client.NewGatewayClient(cfg.Services["gateway"], 1)

	heartBeat := gateway_client.HeartBeat{
		Url: cfg.Info.HeartBeat.Url,
		TTL: int64(cfg.Info.HeartBeat.TTL),
	}

	err = gateway.RegisterService(gateway_client.RegisterRequest{
		Name:      cfg.Info.Name,
		Version:   cfg.Info.Version,
		UrlPrefix: cfg.Info.UrlPrefix,
		Host:      cfg.Server.Host,
		Port:      strconv.Itoa(int(cfg.Server.Port)),
		BaseUrl:   cfg.Info.BaseUrl,
		Mapping: map[string]gateway_client.Endpoint{
			"/trips": {
				Url:            "/trips",
				PermissionList: map[string]any{},
			},
		},
		HeartBeat: heartBeat,
	})

	if err != nil {
		log.Fatal(err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	wg := sync.WaitGroup{}
	wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()

		err = http.Run(appContainer, cfg.Server)

		if err != nil {
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		log.Println("Starting gRPC Server on port 8082")
		err = grpcServer.Serve(l)

		if err != nil {
			return
		}
	}()

	<-ctx.Done()
	fmt.Println("Server received shutdown signal, waiting for components to stop...")
	return

}

func runVehicleRequestsScheduler(a app.App) error {
	ctx := context.Background()
	s := a.VehicleRequestPollerService(ctx)

	err := s.PollAndPublish(ctx, domain.PollerRequest{TotalRecords: 20, BatchSize: 4, ConcurrentJobs: 5})
	if err != nil {
		fmt.Printf("error in running vehicle requests poller: %s\n", err.Error())
		return err

	}

	go func() {
		ticker := time.NewTicker(2 * time.Hour)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				err := s.PollAndPublish(ctx, domain.PollerRequest{TotalRecords: 20, BatchSize: 4, ConcurrentJobs: 5})
				if err != nil {
					fmt.Printf("error in running vehicle requests poller: %s\n", err.Error())
				}

			}
		}
	}()
	return nil

}
