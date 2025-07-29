package server

import (
	"log"
	"os"

	"github.com/OpsOMI/S.L.A.M/internal/configs/server/db"
	"gopkg.in/yaml.v3"
)

const defaultConfigPath = "./configs/server.yaml"

type ServerConfigs struct {
	Host                 string `yaml:"host"`
	Port                 string `yaml:"port"`
	RequireAuth          string `yaml:"require_auth"`
	MessageLifeTimeHours int    `yaml:"message_lifetime_hours"`
	LogLevel             string `yaml:"log_level"`
	MaxClients           int    `yaml:"max_clients"`
	TSLCertPath          string `yaml:"tls_cert_path"`
	TSLKeyPath           string `yaml:"tls_key_path"`
	Env                  db.ENV
}

func LoadConfig(
	path string,
	envFiles ...string,
) *ServerConfigs {
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

	cfg.Env = *db.LoadAll(
		envFiles...,
	)

	return &cfg
}
