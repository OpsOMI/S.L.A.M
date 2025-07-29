package config

import (
	"github.com/OpsOMI/S.L.A.M/internal/server/config/core"
	"github.com/OpsOMI/S.L.A.M/internal/server/config/db"
	"github.com/OpsOMI/S.L.A.M/internal/server/config/managment"
)

type Configs struct {
	Server    core.Configs
	Managment managment.ManagementConfig
	DB        db.DatabaseConfig
}

func LoadConfig(
	serverConfigPath string,
	managmentEnvPath string,
	envFiles ...string,
) *Configs {
	cfg := &Configs{
		Server: *core.LoadConfig(serverConfigPath),
		DB:     *db.LoadAll(envFiles...),
	}

	if managmentEnvPath != "" {
		cfg.Managment = *managment.LoadAll(managmentEnvPath)
	}

	return cfg
}
