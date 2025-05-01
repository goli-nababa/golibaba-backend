package config

type Config struct {
	Server ServerConfig `json:"server"`
	Redis  RedisConfig  `json:"redis"`
	Grpc   GrpcConfig   `json:"grpc"`
}

type ServerConfig struct {
	Port       uint `json:"port"`
	ServiceTTL uint `json:"service_ttl"`
}

type RedisConfig struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}

type GrpcConfig struct {
	Url     string `json:"url"`
	Version uint32 `json:"version"`
	Port    uint64 `json:"port"`
}
