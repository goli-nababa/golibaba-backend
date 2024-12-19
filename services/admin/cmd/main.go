package main

import (
	"admin/api/http"
	"admin/app"
	"admin/config"
	"log"
	"os"
)

func main() {
	path := os.Getenv("CONFIG_FILE")
	if path == "" {
		path = "../config.json"
	}
	cfg := config.MustReadConfig(path)

	appContainer := app.NewMustApp(cfg)

	log.Fatal(http.Run(appContainer, cfg.Server))

}
