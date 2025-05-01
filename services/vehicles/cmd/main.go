package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"vehicles/config"
	"vehicles/dto"
	"vehicles/internal/common/domain"
	"vehicles/pkg/adapters/storage"
	"vehicles/pkg/adapters/storage/migrations"
	"vehicles/pkg/postgres"

	"github.com/goli-nababa/golibaba-backend/modules/trip_service_client"
	proto "github.com/goli-nababa/golibaba-backend/proto/pb"
	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

func main() {
	path := os.Getenv("CONFIG_FILE")
	if path == "" {
		path = "../config.json"
	}

	cfg := config.MustReadConfig(path)

	db, err := getDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = migrations.AutoMigrate(db)
	if err != nil {
		log.Fatal(err)

	}

	vehicleRepo := storage.NewVehicleRepo(db)

	mqConn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.MessageQueue.RabbitMqUsername, cfg.MessageQueue.RabbitMqPassword, cfg.MessageQueue.RabbitMqHost, cfg.MessageQueue.RabbitMqPort))
	if err != nil {
		log.Fatal(err)
	}
	defer mqConn.Close()

	ctx := context.Background()

	runTripVehicleMatcher(ctx, cfg, vehicleRepo, mqConn)

}

func runTripVehicleMatcher(ctx context.Context, cfg config.Config, repo *storage.VehicleRepo, conn *amqp.Connection) {
	// Open a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// Declare a queue (idempotent)
	queueName := cfg.MessageQueue.VehicleRequestQueueName
	q, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}
	// Consume messages from the queue
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer tag
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	forever := make(chan bool)

	go func() {
		filters := []*domain.RepositoryFilter{}
		for d := range msgs {
			var message dto.VehicleRequest
			if err := json.Unmarshal(d.Body, &message); err != nil {
				log.Printf("Error decoding JSON: %v", err)
				continue
			}
			fmt.Printf("Received message: %v\n", message)

			if message.VehicleCost > 0 {
				filters = append(filters, &domain.RepositoryFilter{Field: "cost", Operator: "=", Value: strconv.Itoa(int(message.VehicleCost))})
			}
			if message.VehicleTypeId > 0 {
				filters = append(filters, &domain.RepositoryFilter{Field: "vehicle_type_id", Operator: "=", Value: strconv.Itoa(int(message.VehicleTypeId))})
			}
			if message.VehicleCreationDate != nil {
				filters = append(filters, &domain.RepositoryFilter{Field: "vehicle_creation_date", Operator: "=", Value: message.VehicleCreationDate.String()})
			}
			filters = append(filters, &domain.RepositoryFilter{Field: "passenger_capacity", Operator: ">=", Value: strconv.Itoa(int(message.Trip.PassengersCountLimit))})

			vehicles, err := repo.Get(ctx, &domain.RepositoryRequest{Filters: filters, Limit: 1})
			if err != nil {
				fmt.Println("error in get vehicles: " + err.Error())
			}

			host := cfg.Services["transportation_host"]
			port, _ := strconv.Atoi(cfg.Services["transportation_port"])

			for _, v := range vehicles {
				t, err := trip_service_client.NewTripServiceClient(host, 1, uint64(port))
				if err != nil {
					fmt.Println("error in get trip client: " + err.Error())
				}
				t.SetVehicle(&proto.TripVehicle{TripId: uint64(message.Trip.ID), VehicleId: uint64(v.ID), VehicleSpeed: uint64(v.Speed)})
			}

		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func getDB(cfg config.Config) (*gorm.DB, error) {
	db, err := postgres.NewPsqlGormConnection(postgres.DBConnOptions{
		User:   cfg.DB.User,
		Pass:   cfg.DB.Pass,
		Host:   cfg.DB.Host,
		Port:   cfg.DB.Port,
		Name:   cfg.DB.Name,
		Schema: cfg.DB.Schema,
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
