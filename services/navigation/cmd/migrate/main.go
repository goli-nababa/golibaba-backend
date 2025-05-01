package main

import (
	"flag"
	"log"
	"navigation_service/config"
	"navigation_service/pkg/adapters/storage/migrations"
	"navigation_service/pkg/postgres"
)

func main() {
	configPath := flag.String("config", "config.json", "path to config file")
	flag.Parse()

	// Load configuration
	cfg, err := config.ReadConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	// Connect to database
	db, err := postgres.NewPsqlGormConnection(postgres.DBConnOptions{
		Host:   cfg.DB.Host,
		Port:   cfg.DB.Port,
		User:   cfg.DB.User,
		Pass:   cfg.DB.Password,
		Name:   cfg.DB.Database,
		Schema: cfg.DB.Schema,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations
	migrationManager := migrations.NewManager(db)
	if err := migrationManager.RunMigrations(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Migrations completed successfully")
}
