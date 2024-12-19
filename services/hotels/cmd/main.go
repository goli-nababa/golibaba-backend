package main

import (
	"flag"
	"hotels-service/config"
	"os"
)

var cnfgPath = flag.String("c", "config.json", "to set config")

func main() {
	// config := readConfig()
}

func readConfig() config.Config {
	flag.Parse()
	if envConfigPath := os.Getenv("HOTEL_CONFIG"); len(envConfigPath) > 0 {
		*cnfgPath = envConfigPath
	}
	config := config.MustReadConfig(*cnfgPath)
	return config
}
