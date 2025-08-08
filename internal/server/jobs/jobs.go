package jobs

import (
	"context"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
	"github.com/OpsOMI/S.L.A.M/internal/server/services"
	"github.com/OpsOMI/S.L.A.M/pkg/cronpkg"
)

func Register(
	cronManager *cronpkg.Manager,
	logg logger.ILogger,
	services services.IServices,
) {
	// Example JOB 1: Ping
	cronManager.AddJob("@every 1m", func() {
		logg.Info("Ping!")
	})

	// Delete Every Messages day.
	cronManager.AddJob("@daily", func() {
		if err := services.Messages().DeleteMessages(context.Background()); err != nil {
			logg.Warn("Failed to delete messages: " + err.Error())
		}
	})
}
