package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/goli-nababa/golibaba-backend/common"
)

type GatewayService struct {
	Url     string
	Version int64
}

func NewGatewayClient(url string, version int64) GatewayService {
	return GatewayService{
		Url:     url,
		Version: version,
	}
}

func (gateway *GatewayService) RegisterService(service common.Service) error {
	res, err := http.Get(gateway.Url)

	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)
	return nil
}

func main() {

}
