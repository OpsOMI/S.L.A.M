package core

import (
	"log"
	"os"

	"github.com/OpsOMI/S.L.A.M/internal/server/config/core/models"
	"gopkg.in/yaml.v3"
)

const defaultConfigPath = "./configs/server.yaml"

type Configs struct {
	Core models.Server `yaml:"server"`
	App  models.App    `yaml:"app"`
}

func LoadConfig(
	path string,
) *Configs {
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
