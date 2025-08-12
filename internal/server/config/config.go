package config

import (
	"log"

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

func CheckConfig(configs *Configs) {
	if configs.Env.Managment.Nickname == "" {
		log.Fatal("[config] Management Nickname cannot be empty")
	}
	if configs.Env.Managment.Username == "" {
		log.Fatal("[config] Management Username cannot be empty")
	}
	if configs.Env.Managment.Password == "" {
		log.Fatal("[config] Management Password cannot be empty")
	}

	if configs.Server.Core.Host == "" {
		configs.Server.Core.Host = "localhost"
	}
	if configs.Server.Core.Port == "" {
		configs.Server.Core.Port = "6666"
	}

	if configs.Server.App.Mode == "" {
		log.Fatal("[config] Server mode must be set (dev or prod)")
	}

	switch configs.Server.App.Mode {
	case "dev":
		if configs.Db.Dev.User == "" {
			log.Fatal("[config] DB user cannot be empty in dev mode")
		}
		if configs.Db.Dev.Password == "" {
			log.Fatal("[config] DB password cannot be empty in dev mode")
		}
	case "prod":
		if configs.Db.Prod.User == "" {
			log.Fatal("[config] DB user cannot be empty in prod mode")
		}
		if configs.Db.Prod.Password == "" {
			log.Fatal("[config] DB password cannot be empty in prod mode")
		}

	default:
		log.Fatalf("[config] Unsupported server mode: %s", configs.Server.App.Mode)
	}
}
