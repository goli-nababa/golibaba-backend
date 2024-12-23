package main

import (
	"log"
	"os"
	"transportation/api/http"
	"transportation/app"
	"transportation/config"
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
