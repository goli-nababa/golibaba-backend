package main

import (
	"bank_service/api"
	"bank_service/api/fiber"
	"bank_service/app"
	"bank_service/config"
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	configPath = flag.String("config", "config.json", "path to config file")
)

func main() {
	flag.Parse()

	if envConfig := os.Getenv("CONFIG_PATH"); envConfig != "" {
		*configPath = envConfig
	}

	cfg, err := config.ReadConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	application := app.NewMustApp(cfg)

	go func() {
		if err := api.RunGRPCServer(ctx, cfg); err != nil {
			log.Printf("gRPC server error: %v", err)
		}
	}()

	httpServer := fiber.NewServer(application, cfg)
	go func() {
		if err := httpServer.Start(); err != nil {
			log.Printf("HTTP server error: %v", err)
		}
	}()

	<-ctx.Done()

	if err := httpServer.Shutdown(); err != nil {
		log.Printf("Error shutting down HTTP server: %v", err)
	}

	stop()
	log.Println("Servers gracefully stopped")
}
