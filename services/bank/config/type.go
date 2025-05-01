package config

type Config struct {
	DB     DBConfig     `json:"db"`
	Server ServerConfig `json:"server"`
	Redis  RedisConfig  `json:"redis"`
	Logger LoggerConfig `json:"logger"`
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
	HTTPPort              uint   `json:"http_port"`
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
