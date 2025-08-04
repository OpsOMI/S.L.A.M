package app

import (
	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
	"github.com/OpsOMI/S.L.A.M/internal/client/config"
	"github.com/OpsOMI/S.L.A.M/internal/client/infrastructure/network"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/controller"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/response"
)

func Run(cfg *config.Configs) {
	logg, err := logger.NewZapLogger()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	defer logg.Sync()

	logg.Info("Welcome to the S.L.A.M client!")

	logg.Info("Attempting to connect to server...")
	conn, err := network.ConnectToServer(
		cfg.ServerName,
		cfg.ServerHost,
		cfg.ServerPort,
		cfg.TSLCertPath,
		cfg.TimeoutSeconds,
		cfg.ReconnectRetry,
	)
	if err != nil {
		logg.Error("Failed to connect to server: " + err.Error())
		return
	}
	defer conn.Close()
	logg.Info("Successfully connected to server")

	// This is for now. This reads the first welcome to slam message.
	response.Read(conn)

	logg.Info("Controller Started")
	controller := controller.NewController(conn, logg)
	controller.Run()
}
