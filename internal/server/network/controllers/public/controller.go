package public

import (
	"encoding/json"
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/commons"
	"github.com/OpsOMI/S.L.A.M/internal/server/infrastructure/connection"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/response"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/store"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/types"
	"github.com/OpsOMI/S.L.A.M/internal/server/services"
)

type Controller struct {
	logger     logger.ILogger
	routes     map[string]types.HandlerFunc
	store      store.IJwtManager
	services   services.IServices
	connecions *connection.ConnectionManager
}

func NewController(
	logger logger.ILogger,
	store store.IJwtManager,
	services services.IServices,
	connections *connection.ConnectionManager,
) *Controller {
	pc := &Controller{
		logger:     logger,
		routes:     make(map[string]types.HandlerFunc),
		store:      store,
		services:   services,
		connecions: connections,
	}

	pc.InitHealthRoutes()
	pc.InitAuthRoutes()

	return pc
}

func (p *Controller) Route(
	conn net.Conn,
	cmd string,
	args json.RawMessage,
) error {
	handler, ok := p.routes[cmd]
	if !ok {
		return response.Response(commons.StatusBadRequest, "Unknown Command", cmd)
	}

	return handler(conn, args, nil)
}
