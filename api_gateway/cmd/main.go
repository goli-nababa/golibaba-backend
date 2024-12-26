package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/goli-nababa/golibaba-backend/modules/cache"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"api_gateway/api/gateway"
	"api_gateway/api/http/types"
	"api_gateway/app"
	"api_gateway/config"

	httpHandler "api_gateway/api/http"
)

var configPath = flag.String("config", "config.json", "Path to service config file")

func main() {
	flag.Parse()

	if v := os.Getenv("CONFIG_PATH"); len(v) > 0 {
		*configPath = v
	}

	c := config.MustReadConfig(*configPath)

	appContainer := app.MustNewApp(c)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	wg := sync.WaitGroup{}
	wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := httpHandler.Bootstrap(appContainer, c.Server)

		if err != nil {
			return
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		cacheProvider := cache.NewJsonObjectCache[*types.RegisterRequest](
			appContainer.Cache(),
			"service.",
		)

		handler := gateway.NewGateway(cacheProvider)

		log.Println("Starting API Gateway on :8081")
		if err := http.ListenAndServe(":8081", handler); err != nil {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	<-ctx.Done()
	fmt.Println("Server received shutdown signal, waiting for components to stop...")
	return
}
