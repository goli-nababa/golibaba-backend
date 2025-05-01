package adapters

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQPublisher struct {
	Channel *amqp.Channel
	Queue   string
}

func NewRabbitMQPublisher(channel *amqp.Channel, queue string) *RabbitMQPublisher {
	return &RabbitMQPublisher{
		Channel: channel,
		Queue:   queue,
	}
}

func (p *RabbitMQPublisher) Publish(message []byte) error {
	err := p.Channel.Publish(
		"",      // exchange
		p.Queue, // routing key
		false,   // mandatory
		false,   // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        message,
		},
	)
	if err != nil {
		log.Printf("Failed to publish message: %v", err)
		return err
	}

	return nil
}
