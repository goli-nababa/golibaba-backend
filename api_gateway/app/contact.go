package app

import (
	"api_gateway/config"
	adapters "api_gateway/pkg/adapters/rabbitmq"

	"github.com/goli-nababa/golibaba-backend/modules/cache"
)

type App interface {
	Config() config.Config
	Cache() cache.Provider
	RabbitMQPublisher() *adapters.RabbitMQPublisher
	StartRabbitMQConsumer(consumerFunc func(channel *adapters.RabbitMQConsumer))
}
