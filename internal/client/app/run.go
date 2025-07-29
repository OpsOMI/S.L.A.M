package app

import (
	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
	"github.com/OpsOMI/S.L.A.M/internal/adapters/network"
	"github.com/OpsOMI/S.L.A.M/internal/client/config"
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

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		logg.Error("Failed to read message from server: " + err.Error())
		return
	}

	message := string(buf[:n])
	logg.Info("Received message from server: " + message)
}
