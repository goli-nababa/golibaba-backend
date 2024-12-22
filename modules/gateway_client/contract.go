package main

import (
	"github.com/goli-nababa/golibaba-backend/common"
)

type GatewayService interface {
	RegisterService(service common.Service) error
}
