package app

import "api_gateway/config"

type App interface {
	Config() config.Config
}
