package app

import (
	"os"
	"strings"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
	"github.com/OpsOMI/S.L.A.M/internal/client/config"
	"github.com/OpsOMI/S.L.A.M/internal/client/infrastructure/network"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/controller"
)

func Run(
	cfg *config.Configs,
) {
	logg, err := logger.NewZapLogger()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	defer logg.Sync()

	logg.Info("Welcome to the S.L.A.M client!")

	logg.Info("Attempting to connect to server...")
	var certData []byte
	if strings.EqualFold(cfg.UseEmbed, "true") {
		certData = config.EmbededTSKCertBinary
	} else {
		certData, err = os.ReadFile(cfg.TSLCertPath)
		if err != nil {
			logg.Error("Failed to read certificate: " + err.Error())
			return
		}
	}

	conn, err := network.ConnectToServer(
		cfg.ServerName,
		cfg.ServerHost,
		cfg.ServerPort,
		certData,
		cfg.TimeoutSeconds,
		cfg.ReconnectRetry,
	)
	if err != nil {
		logg.Error("Failed to connect to server: " + err.Error())
		return
	}
	defer conn.Close()
	logg.Info("Successfully connected to server")

	logg.Info("Controller Started")
	controller := controller.NewController(conn, logg, *cfg)
	controller.Run()
}
