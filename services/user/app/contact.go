package app

import "user_service/config"

type App interface {
	Config() config.Config
}
