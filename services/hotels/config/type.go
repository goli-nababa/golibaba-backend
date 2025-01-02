package config

type Config struct {
	Database       DBConfig          `json:"db"`
	Server         ServerConfig      `json:"server"`
	Redis          RedisConfig       `json:"redis"`
	Logger         LoggerConfig      `json:"logger"`
	Info           ServiceInfo       `json:"service_info"`
	Services       map[string]string `json:"services"`
	UserService    UserServiceConfig `json:"user_service"`
	ConnectGateway bool              `json:"connect_gateway"`
}

type LoggerConfig struct {
	FilePath string `json:"file_path"`
	Encoding string `json:"encoding"`
	Level    string `json:"level"`
	Logger   string `json:"logger"`
}

type DBConfig struct {
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	Database string `json:"database"`
	Schema   string `json:"schema"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type ServerConfig struct {
	Port                  uint   `json:"port"`
	Host                  string `json:"host"`
	Secret                string `json:"secret"`
	PasswordSecret        string `json:"password_secret"`
	OtpTtlMinutes         uint   `json:"otp_ttl_minutes"`
	MaxRequestsPerSecond  uint   `json:"maxRequestsPerSecond"`
	AuthExpirationMinutes uint   `json:"auth_expiration_minutes"`
	AuthRefreshMinutes    uint   `json:"auth_refresh_minutes"`
}

type RedisConfig struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}

type HeatBeat struct {
	Url string `json:"url"`
	TTL uint   `json:"ttl"`
}

type ServiceInfo struct {
	Name      string   `json:"name"`
	Version   string   `json:"version"`
	UrlPrefix string   `json:"url_prefix"`
	BaseUrl   string   `json:"base_url"`
	HeartBeat HeatBeat `json:"heart_beat"`
}

type UserServiceConfig struct {
	URL     string `json:"url"`
	Port    uint   `json:"port"`
	Version uint32 `json:"version"`
}
