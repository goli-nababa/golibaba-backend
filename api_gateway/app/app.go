package app

import (
	"api_gateway/config"
	adapters "api_gateway/pkg/adapters/rabbitmq"
	"fmt"

	user "github.com/goli-nababa/golibaba-backend/modules/user_service_client"

	"github.com/goli-nababa/golibaba-backend/modules/cache"
)

type app struct {
	cfg          config.Config
	redis        cache.Provider
	rabbitMQConn *adapters.RabbitMQConnection
	publisher    *adapters.RabbitMQPublisher
}

func (a *app) Config() config.Config {
	return a.cfg
}

func (a *app) Cache() cache.Provider {
	return a.redis
}

func (a *app) RabbitMQPublisher() *adapters.RabbitMQPublisher {
	return a.publisher
}

func (a *app) StartRabbitMQConsumer(consumerFunc func(channel *adapters.RabbitMQConsumer)) {
	if a.rabbitMQConn == nil {
		channel, err := adapters.InitializeRabbitMQConnection()
		if err != nil {
			panic(fmt.Errorf("failed to initialize RabbitMQ connection: %v", err))
		}
		a.rabbitMQConn = channel
	}

	userServiceClient, _ := user.NewUserServiceClient(a.cfg.Grpc.Url, a.cfg.Grpc.Version, a.cfg.Grpc.Port)
	consumer := adapters.NewRabbitMQConsumer(a.rabbitMQConn.Channel, "logs_queue", userServiceClient)
	go consumerFunc(consumer)
}

func (a *app) setRabbitMQ() {
	channel, err := adapters.InitializeRabbitMQConnection()
	if err != nil {
		panic(fmt.Errorf("failed to initialize RabbitMQ connection: %v", err))
	}

	a.rabbitMQConn = channel
	a.publisher = adapters.NewRabbitMQPublisher(channel.Channel, "logs_queue")
}

func (a *app) setRedis() {
	a.redis = cache.NewRedisProvider(
		fmt.Sprintf("%s:%d", a.cfg.Redis.Host, a.cfg.Redis.Port),
		"", "", 0,
	)
}

func NewApp(cfg config.Config) (App, error) {
	a := &app{cfg: cfg}

	/*	if err := a.setDB(); err != nil {
		return nil, err
	}*/

	a.setRedis()

	return a, nil
}

func MustNewApp(cfg config.Config) App {
	a, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return a
}
