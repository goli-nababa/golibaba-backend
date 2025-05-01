package adapters

import (
	"github.com/streadway/amqp"
)

type RabbitMQConnection struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func InitializeRabbitMQConnection() (*RabbitMQConnection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@golibaba-rabbitmq/")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQConnection{Connection: conn, Channel: ch}, nil
}
