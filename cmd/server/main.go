package main

import (
	"github.com/OpsOMI/S.L.A.M/internal/server/app"
	"github.com/OpsOMI/S.L.A.M/internal/server/config"
)

func main() {
	configs := config.LoadConfig(
		"./configs/server.yaml",
		"./env/real/.env.management",
		"./deployment/dev/.env",
		"./deployment/prod/.env",
	)
	app.Run(*configs)
}
