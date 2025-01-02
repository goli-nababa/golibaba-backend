package config

type Config struct {
	DB       DBConfig          `json:"db"`
	Server   ServerConfig      `json:"server"`
	Redis    RedisConfig       `json:"redis"`
	Info     ServiceInfo       `json:"service_info"`
	Services map[string]string `json:"services"`
	SMTP     SMTPConfig        `json:"smtp"`
}

type ServiceInfo struct {
	Name      string   `json:"name"`
	Version   string   `json:"version"`
	UrlPrefix string   `json:"url_prefix"`
	BaseUrl   string   `json:"base_url"`
	HeartBeat HeatBeat `json:"heart_beat"`
}

type HeatBeat struct {
	Url string `json:"url"`
	TTL uint   `json:"ttl"`
}

type SMTPConfig struct {
	Email    string `json:"email"`
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type DBConfig struct {
	Host   string `json:"host"`
	Port   uint   `json:"port"`
	User   string `json:"user"`
	Pass   string `json:"pass"`
	Name   string `json:"name"`
	Schema string `json:"schema"`
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
