package app

import (
	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
	"github.com/OpsOMI/S.L.A.M/internal/adapters/postgres"
	"github.com/OpsOMI/S.L.A.M/internal/adapters/postgres/sqlc/pgqueries"
	"github.com/OpsOMI/S.L.A.M/internal/server/config"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains"
	"github.com/OpsOMI/S.L.A.M/internal/server/infrastructure/network"
	"github.com/OpsOMI/S.L.A.M/internal/server/jobs"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/controllers"
	"github.com/OpsOMI/S.L.A.M/internal/server/repositories"
	"github.com/OpsOMI/S.L.A.M/internal/server/services"
	"github.com/OpsOMI/S.L.A.M/pkg"
	"github.com/OpsOMI/S.L.A.M/pkg/cronpkg"
	"go.uber.org/zap"
)

func Run(cfg config.Configs) {
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

	// Initialize domain mapper
	mappers := domains.NewMappers()

	// Initialize sqlc queries with DB connection and mapper
	queries := pgqueries.New(conn)

	// Initialize common/shared packages (Hasher, Mailer, etc.)
	packages := pkg.NewPackages(
		conn,
	)
	logg.Info("Shared packages initialized")

	// Initialize repositories with queries
	repositories := repositories.NewRepositories(queries, mappers, packages.TXManager())
	logg.Info("Repositories initialized")

	services := services.NewServices(logg, packages, repositories)
	logg.Info("Services initialized")

	Setup(cfg, services, logg)
	logg.Info("[setup] Default migartions added successfully.")

	// Initialize cron job manager and register jobs
	cronManager := cronpkg.New()
	jobs.Register(cronManager, logg, services)
	cronManager.Start()
	logg.Info("Cron jobs started")

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

	logg.Info("Controller Starting...")
	controller := controllers.NewController(listener, logg, cfg)

	if err := controller.Start(services, &cfg); err != nil {
		logg.Error("Controller stopped with error", zap.Error(err))
	}
}
