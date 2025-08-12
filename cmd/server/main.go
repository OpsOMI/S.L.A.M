package main

import (
	"os"

	"github.com/OpsOMI/S.L.A.M/internal/server/app"
	"github.com/OpsOMI/S.L.A.M/internal/server/config"
)

func main() {
	mode := os.Getenv("APP_MODE")
	if mode == "" {
		mode = "dev"
	}

	configs := config.LoadConfig(
		mode,
		"./configs/server.yaml",
		"./env/real/.env",
		"./deployment/dev/.env",
		"./deployment/prod/.env",
	)

	app.Run(*configs)
}
