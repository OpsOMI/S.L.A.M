package server

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const defaultConfigPath = "./configs/server.yaml"

type ServerConfigs struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func LoadConfig(path string) *ServerConfigs {
	if path == "" {
		path = defaultConfigPath
	}

	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}

	var cfg ServerConfigs
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		log.Fatalf("failed to unmarshal config: %v", err)
	}

	return &cfg
}
