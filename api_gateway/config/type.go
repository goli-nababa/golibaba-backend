package config

type Config struct {
	Server ServerConfig `json:"server"`
	Redis  RedisConfig  `json:"redis"`
}

type ServerConfig struct {
	Port       uint `json:"port"`
	ServiceTTL uint `json:"service_ttl"`
}

type RedisConfig struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}
