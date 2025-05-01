package config

type Config struct {
	DB           DBConfig           `json:"db"`
	Server       ServerConfig       `json:"server"`
	Redis        RedisConfig        `json:"redis"`
	Logging      LoggerConfig       `json:"logging"`
	Info         ServiceInfo        `json:"service_info"`
	MessageQueue MessageQueueConfig `json:"message_queue"`
	Services     map[string]string  `json:"services"`
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

type MessageQueueConfig struct {
	RabbitMqHost     string `json:"rabbit_mq_host"`
	RabbitMqPort     string `json:"rabbit_mq_port"`
	RabbitMqUsername string `json:"rabbit_mq_username"`
	RabbitMqPassword string `json:"rabbit_mq_password"`

	VehicleRequestQueueName string `json:"vehicle_request_queue_name"`
}
