package config

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
)

func absPath(path string) (string, error) {
	if filepath.IsAbs(path) {
		return path, nil
	}
	return filepath.Abs(path)
}
func ReadConfig[T any](configPath string) (T, error) {
	var config T

	fullPath, err := absPath(configPath)
	if err != nil {
		return config, err
	}
	viper.SetConfigFile(fullPath)
	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return config, err
	}

	return config, viper.Unmarshal(&config)
}

func MustReadConfig(configPath string) Config {
	config, err := ReadConfig[Config](configPath)
	if err != nil {
		panic(fmt.Errorf("failed to read config :%w", err))
	}
	return config
}
