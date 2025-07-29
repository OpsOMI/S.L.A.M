package runserver

import (
	"log"
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
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

	ln, err := net.Listen("tcp", "localhost:6666")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	logg.Info("Server Listening")
}
