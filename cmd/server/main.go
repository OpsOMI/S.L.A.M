package main

import (
	"os"

	"github.com/OpsOMI/S.L.A.M/internal/server/app"
	"github.com/OpsOMI/S.L.A.M/internal/server/config"
)

func main() {
	mode := os.Getenv("APP_MODE")
	if mode == "" {
		mode = "dexv"
	}

	configs := config.LoadConfig(
		"./configs/server.yaml",
		"./env/real/.env",
		"./deployment/dev/.env",
		"./deployment/prod/.env",
	)
	configs.Server.App.Mode = mode

	app.Run(*configs)
}
