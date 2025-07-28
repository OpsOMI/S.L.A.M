package client

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const defaultConfigPath = "./configs/client.yaml"

type ClientConfigs struct {
	Nickname string `yaml:"nickname"`
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
