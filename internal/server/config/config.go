package config

import (
	"github.com/OpsOMI/S.L.A.M/internal/server/config/core"
	"github.com/OpsOMI/S.L.A.M/internal/server/config/db"
	"github.com/OpsOMI/S.L.A.M/internal/server/config/env"
)

type Configs struct {
	Server core.Configs
	Env    env.EnvConfig
	Db     db.DatabaseConfig
}

func LoadConfig(
	serverConfigPath string,
	managmentEnvPath string,
	envFiles ...string,
) *Configs {
	cfg := &Configs{
		Server: *core.LoadConfig(serverConfigPath),
		Db:     *db.LoadAll(envFiles...),
	}

	if managmentEnvPath != "" {
		cfg.Env = *env.LoadAll(managmentEnvPath)
	}

	return cfg
}
