package server

import (
	"github.com/OpsOMI/S.L.A.M/internal/configs/server/core"
	"github.com/OpsOMI/S.L.A.M/internal/configs/server/db"
	"github.com/OpsOMI/S.L.A.M/internal/configs/server/managment"
)

type ServerConfigs struct {
	Server    core.Configs
	Managment managment.ManagementConfig
	DB        db.DatabaseConfig
}

func LoadConfig(
	serverConfigPath string,
	managmentEnvPath string,
	envFiles ...string,
) *ServerConfigs {
	cfg := &ServerConfigs{
		Server: *core.LoadConfig(serverConfigPath),
		DB:     *db.LoadAll(envFiles...),
	}

	if managmentEnvPath != "" {
		cfg.Managment = *managment.LoadAll(managmentEnvPath)
	}

	return cfg
}
