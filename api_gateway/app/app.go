package app

import "api_gateway/config"

type app struct {
	cfg config.Config
}

func (a *app) Config() config.Config {
	return a.cfg
}

func NewApp(cfg config.Config) (App, error) {
	a := &app{cfg: cfg}

	/*	if err := a.setDB(); err != nil {
			return nil, err
		}

		a.setRedis()
		a.setEmailService()*/

	return a, nil
}

func MustNewApp(cfg config.Config) App {
	a, err := NewApp(cfg)
	if err != nil {
		panic(err)
	}
	return a
}
