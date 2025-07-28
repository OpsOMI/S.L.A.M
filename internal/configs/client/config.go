package client

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const defaultConfigPath = "./configs/client.yaml"

type ClientConfigs struct {
	ServerHost     string `yaml:"server_host"`
	ServerPort     string `yaml:"server_port"`
	TimeoutSeconds int    `yaml:"timeout_seconds"`
	ReconnectRetry int    `yaml:"reconnect_retry"`
}

func LoadConfig(path string) *ClientConfigs {
	if path == "" {
		path = defaultConfigPath
	}

	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	var cfg ClientConfigs
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		log.Fatalf("failed to unmarshal config: %v", err)
	}

	return &cfg
}
