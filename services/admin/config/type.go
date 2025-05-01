package config

type Config struct {
	DB      DBConfig     `json:"db"`
	Server  ServerConfig `json:"server"`
	Redis   RedisConfig  `json:"redis"`
	Logging LoggerConfig `json:"logging"`
}

type DBConfig struct {
	Host   string `json:"host"`
	Port   uint   `json:"port"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
	Name   string `json:"name"`
	Schema string `json:"schema"`
}

type LoggerConfig struct {
	FilePath string `json:"filePath"`
	Encoding string `json:"encoding"`
	Level    string `json:"level"`
	Logger   string `json:"logger"`
}

type ServerConfig struct {
	Port                  uint   `json:"port"`
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

type GrpcConfig struct {
	Url     string `json:"url"`
	Version uint32 `json:"version"`
	Port    uint64 `json:"port"`
}
