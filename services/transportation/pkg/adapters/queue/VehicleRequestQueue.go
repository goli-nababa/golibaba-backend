package queue

import (
	"context"
	"encoding/json"
	"transportation/internal/trip/domain"

	"github.com/streadway/amqp"
)

type VehicleRequestQueueRepo struct {
	mqConn *amqp.Connection
	queue  string
}

func NewVehicleRequestQueueRepo(mqConn *amqp.Connection, queue string) *VehicleRequestQueueRepo {
	return &VehicleRequestQueueRepo{mqConn: mqConn, queue: queue}
}

func (p *VehicleRequestQueueRepo) PublishRequest(ctx context.Context, msg domain.VehicleRequest) error {

	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	channel, err := p.mqConn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	err = channel.Publish(
		"",
		p.queue,
		true,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        jsonMsg,
		},
	)

	if err != nil {
		return err
	}

	return nil
}
