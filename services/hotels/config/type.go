package config

type Config struct {
	Database Postgresql `mapstructure:"postgresql"`
}

type Postgresql struct {
	Host     string `mapstructure:"host"`
	Port     uint   `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	SSLMode  string `mapstructure:"sslmode"`
	Schema   string `mapstructure:"schema"`
}
