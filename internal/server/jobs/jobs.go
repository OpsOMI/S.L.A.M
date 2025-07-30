package jobs

import (
	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
	"github.com/OpsOMI/S.L.A.M/pkg/cronpkg"
)

func Register(
	cronManager *cronpkg.Manager,
	logg logger.ILogger,
) {
	// Example JOB 1: Ping
	cronManager.AddJob("@every 1m", func() {
		logg.Info("Ping!")
	})

	// Example JOB 2: API key expire
	// cronManager.AddJob("0 0 * * *", func() {
	// 	services.APIKey().ExpireOldKeys()
	// 	logg.Info("API keys expired.")
	// })
}
