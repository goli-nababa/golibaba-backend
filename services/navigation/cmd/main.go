package main

import (
	"flag"
	"fmt"
	"log"
	"navigation_service/api/grpc_server"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"navigation_service/api/http"
	"navigation_service/app"
	"navigation_service/config"
)

var configPath = flag.String("config", "config.json", "path to config file")

func main() {
	flag.Parse()

	if envConfig := os.Getenv("CONFIG_PATH"); envConfig != "" {
		*configPath = envConfig
	}

	cfg, err := config.ReadConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	newApp, err := app.NewApp(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup
	errChan := make(chan error, 2)

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Printf("Starting HTTP server on port %d...\n", cfg.Server.Port)
		if err := http.Bootstrap(newApp, cfg); err != nil {
			errChan <- fmt.Errorf("HTTP server error: %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		grpcPort := cfg.Server.Port + 1
		fmt.Printf("Starting gRPC server on port %d...\n", grpcPort)
		if err := grpc_server.Bootstrap(newApp, grpcPort); err != nil {
			errChan <- fmt.Errorf("gRPC server error: %v", err)
		}
	}()

	select {
	case err := <-errChan:
		log.Printf("Server error: %v", err)
	case sig := <-sigChan:
		log.Printf("Received signal: %v", sig)
	}

	log.Println("Shutting down servers...")

	//wg.Wait()
	log.Println("All servers stopped")
}
