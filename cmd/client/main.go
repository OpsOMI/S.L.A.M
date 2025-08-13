package main

import (
	"github.com/OpsOMI/S.L.A.M/internal/client/app"
	"github.com/OpsOMI/S.L.A.M/internal/client/config"
)

var useEmbed string

func main() {
	var cfg *config.Configs
	if useEmbed == "true" {
		cfg = config.LoadEmbeddedConfig()
	} else {
		cfg = config.LoadConfig("./configs/client.yaml")
	}
	cfg.UseEmbed = useEmbed

	app.Run(cfg)
}
