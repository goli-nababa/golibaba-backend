package gateway_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type gatewayService struct {
	Url     string
	Version uint64
}

func NewGatewayClient(url string, version uint64) GatewayService {
	return &gatewayService{
		Url:     url,
		Version: version,
	}
}

func (gateway *gatewayService) RegisterService(service RegisterRequest) error {
	url := fmt.Sprintf("%s/v%d/register", gateway.Url, gateway.Version)

	marshal, err := json.Marshal(service)

	if err != nil {
		return err
	}

	res, err := http.Post(url, "application/json", bytes.NewReader(marshal))

	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	if res.StatusCode != 201 {
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("client: could not read response body: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("client: status code: %d\n", res.StatusCode)
		fmt.Printf("client: body: %v\n", string(resBody))
	}

	return nil
}
