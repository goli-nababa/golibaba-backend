package common

type HeartBeat struct {
	Url string
	TTL int64
}

type Endpoint struct {
	Url            string
	PermissionList []string // ToDo Implement permission domain
}

// Service represents a microservice configuration and runtime metadata.
// It includes details like the service name, version, routing information,
// inter-service communication settings, and default headers for requests.
type Service struct {
	Name      string              // The name of the service.
	Version   string              // The version of the service.
	UrlPrefix string              // The URL prefix for routing requests to this service.
	BaseUrl   string              // The base URL used for inter-service communication.
	Mapping   map[string]Endpoint // Internal mapping of endpoints for the service.
	HeartBeat HeartBeat           // Heartbeat configuration for monitoring the service's health.
	Headers   map[string]any      // Default headers to be applied to outgoing requests.
}
