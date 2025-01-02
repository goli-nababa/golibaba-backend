package logging

import (
	adapters "api_gateway/pkg/adapters/rabbitmq"
	"api_gateway/pkg/logging/ports"
	"log"
)

type RabbitMQLogPublisher struct {
	Publisher *adapters.RabbitMQPublisher
}

func NewRabbitMQLogPublisher(publisher *adapters.RabbitMQPublisher) ports.LogPublisher {
	return &RabbitMQLogPublisher{Publisher: publisher}
}

func (r *RabbitMQLogPublisher) PublishLog(message []byte) error {
	err := r.Publisher.Publish(message)
	if err != nil {
		log.Printf("Failed to publish log message: %v", err)
		return err
	}
	return nil
}
