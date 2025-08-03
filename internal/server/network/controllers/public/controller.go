package public

import (
	"encoding/json"
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/response"
	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/tokenstore"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/commons"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/types"
	"github.com/OpsOMI/S.L.A.M/internal/server/services"
)

type Controller struct {
	logger     logger.ILogger
	routes     map[string]types.HandlerFunc
	tokenstore tokenstore.ITokenStore
	services   services.IServices
}

func NewController(
	logger logger.ILogger,
	tokenstore tokenstore.ITokenStore,
	services services.IServices,
) *Controller {
	pc := &Controller{
		logger:     logger,
		routes:     make(map[string]types.HandlerFunc),
		tokenstore: tokenstore,
		services:   services,
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
