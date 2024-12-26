package app

import (
	"api_gateway/config"
	"fmt"
	"github.com/goli-nababa/golibaba-backend/modules/cache"
)

type app struct {
	cfg   config.Config
	redis cache.Provider
}

func (a *app) Config() config.Config {
	return a.cfg
}

func (a *app) Cache() cache.Provider {
	return a.redis
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
