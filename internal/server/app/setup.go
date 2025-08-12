package app

import (
	"context"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
	"github.com/OpsOMI/S.L.A.M/internal/server/config"
	"github.com/OpsOMI/S.L.A.M/internal/server/services"
)

func Setup(
	cfg config.Configs,
	svcs services.IServices,
	log logger.ILogger,
) {
	// Create Admin
	SetupDefaultAdmin(cfg, svcs, log)

}

func SetupDefaultAdmin(
	cfg config.Configs,
	svcs services.IServices,
	log logger.ILogger,
) {
	ctx := context.Background()

	ok, err := svcs.Users().IsExistsByUsername(ctx, cfg.Env.Managment.Username)
	if err != nil {
		log.Warnf("[setup] Default owner is_exist failed: %v", err)
		return
	}
	if ok {
		log.Info("[setup] Default owner already exists, skipping creation.")
		return
	}

	id, clientKey, err := svcs.Users().CreateUser(ctx, "cetinboran", cfg.Env.Managment.Username, cfg.Env.Managment.Password, "owner")
	if err != nil {
		log.Warnf("[setup] Default owner creation failed: %v", err)
		return
	}
	log.Infof("[setup] Default owner user created successfully. ID: %v", id)

	if err := svcs.Clients().CreateClient(&cfg, *clientKey); err != nil {
		log.Warnf("[setup] Creating client failed: %v", err)
		return
	}

	roomID, err := svcs.Rooms().Create(ctx, id.String(), "room")
	if err != nil {
		log.Warnf("[setup] Default room creation failed: %v", err)
		return
	}
	log.Infof("[setup] Default room created successfully. ID: %v", roomID)

	emptyrom, _ := svcs.Rooms().Create(ctx, id.String(), "")
	log.Infof("[setup] Default room created successfully. ID: %v", emptyrom)
}
