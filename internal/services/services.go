package services

import (
	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
	"github.com/OpsOMI/S.L.A.M/internal/configs/server"
)

type IServices interface{}

type services struct {
	cfg    server.ServerConfigs
	logger logger.ILogger
}

func NewServices(
	cfg server.ServerConfigs,
	logger logger.ILogger,
) IServices {
	return &services{
		cfg:    cfg,
		logger: logger,
	}
}
