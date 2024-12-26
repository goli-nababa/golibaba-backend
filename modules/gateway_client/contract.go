package gateway_client

type GatewayService interface {
	RegisterService(service RegisterRequest) error
}

type RegisterRequest struct {
	Name      string              `json:"name" validate:"required,min=4"`
	Port      string              `json:"port" validate:"required,min=4"`
	Host      string              `json:"host" validate:"required,min=4"`
	Version   string              `json:"version" validate:"required,min=1,semver"`
	UrlPrefix string              `json:"url_prefix" validate:"required"`
	BaseUrl   string              `json:"base_url" validate:"required"`
	Mapping   map[string]Endpoint `json:"mapping" validate:"required"`
	HeartBeat HeartBeat           `json:"heart_beat" validate:""`
	Headers   map[string]any      `json:"headers" validate:""`
}

type HeartBeat struct {
	Url string `json:"url" validate:""`
	TTL int64  `json:"ttl" validate:""`
}

type Endpoint struct {
	Url            string         `json:"url" validate:"required,min=4"`
	PermissionList map[string]any `json:"permission_list" validate:"required"`
}
