package app

import (
	"api_gateway/config"
	"github.com/goli-nababa/golibaba-backend/modules/cache"
)

type App interface {
	Config() config.Config
	Cache() cache.Provider
}
