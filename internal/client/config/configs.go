package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const defaultConfigPath = "./configs/client.yaml"

type Configs struct {
	ClientKey      string `yaml:"client_key"`
	ServerName     string `yaml:"server_name"`
	ServerHost     string `yaml:"server_host"`
	ServerPort     string `yaml:"server_port"`
	TSLCertPath    string `yaml:"tls_cert_path"`
	TimeoutSeconds int    `yaml:"timeout_seconds"`
	ReconnectRetry int    `yaml:"reconnect_retry"`
	UseEmbed       string
}

// LoadConfig loads config from a file
func LoadConfig(path string) *Configs {
	if path == "" {
		path = defaultConfigPath
	}

	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	var cfg Configs
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		log.Fatalf("failed to unmarshal config: %v", err)
	}

	return &cfg
}
