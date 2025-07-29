package services

import (
	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
	"github.com/OpsOMI/S.L.A.M/internal/server/config"
)

type IServices interface{}

type services struct {
	cfg    config.Configs
	logger logger.ILogger
}

func NewServices(
	cfg config.Configs,
	logger logger.ILogger,
) IServices {
	return &services{
		cfg:    cfg,
		logger: logger,
	}
}
