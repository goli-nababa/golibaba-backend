package adapters

import (
	"encoding/json"
	"log"

	"github.com/goli-nababa/golibaba-backend/common" // Import log domain from common
	user "github.com/goli-nababa/golibaba-backend/modules/user_service_client"
	pb "github.com/goli-nababa/golibaba-backend/proto/pb"
	"github.com/streadway/amqp"
)

type RabbitMQConsumer struct {
	Channel           *amqp.Channel
	Queue             string
	UserServiceClient user.UserServiceClient
}

func NewRabbitMQConsumer(channel *amqp.Channel, queue string, usc user.UserServiceClient) *RabbitMQConsumer {
	return &RabbitMQConsumer{
		Channel:           channel,
		Queue:             queue,
		UserServiceClient: usc,
	}
}

func (r *RabbitMQConsumer) Consume() error {
	msgs, err := r.Channel.Consume(
		r.Queue, // Queue name
		"",      // Consumer tag
		true,    // Auto-acknowledge
		false,   // Exclusive
		false,   // No-local
		false,   // No-wait
		nil,     // Arguments
	)
	if err != nil {
		log.Fatalf("Failed to start consuming: %v", err)
		return err
	}

	for msg := range msgs {
		var logData common.Log
		if err := json.Unmarshal(msg.Body, &logData); err != nil {
			log.Printf("Error unmarshalling log data: %v", err)
			continue
		}

		pbLog := &pb.Log{
			Id:        uint64(logData.ID),
			UserId:    uint64(logData.UserID),
			CompanyId: uint64(logData.CompanyID),
			Action:    logData.Action,
			Path:      logData.Path,
		}

		response, err := r.UserServiceClient.SaveLog(pbLog)
		if err != nil {
			log.Printf("Error saving log via user service client: %v", err)
			continue
		}

		log.Printf("Log saved successfully: %v", response)
	}

	return nil
}
