package runserver

import (
	"fmt"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
	"github.com/OpsOMI/S.L.A.M/internal/adapters/network"
	"github.com/OpsOMI/S.L.A.M/internal/adapters/postgres"
	"github.com/OpsOMI/S.L.A.M/internal/configs/server"
	"go.uber.org/zap"
)

func Run(cfg server.ServerConfigs) {
	logg, err := logger.NewZapLogger()
	if err != nil {
		panic("Failed to initialize logger: " + err.Error())
	}
	defer logg.Sync()

	conn, err := postgres.Connect(cfg)
	if err != nil {
		logg.Error("Failed to connect to DB", zap.Error(err))
		panic(err)
	}
	defer conn.Close()

	if err := postgres.Migrate(conn, cfg.Server.App.MigrationPath); err != nil {
		logg.Error("Migration failed", zap.Error(err))
		panic(err)
	}
	logg.Info("Database connected and migrations applied successfully")

	logg.Info("Server Starting...")
	listener, err := network.StartServer(
		cfg.Server.App.Mode,
		cfg.Server.Core.Port,
		cfg.Server.Core.TSLCertPath,
		cfg.Server.Core.TSLKeyPath,
	)
	if err != nil {
		logg.Errorf("Start Server Error %s", err)
		panic(err)
	}
	logg.Infof("Server Listening on %s", cfg.Server.Core.Port)

	fmt.Println(listener)
}
