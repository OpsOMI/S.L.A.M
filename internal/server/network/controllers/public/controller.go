package public

import (
	"encoding/json"
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/response"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/commons"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/types"
)

type PublicController struct {
	logger logger.ILogger
	routes map[string]types.HandlerFunc
}

func NewController(
	logger logger.ILogger,
) *PublicController {
	pc := &PublicController{
		logger: logger,
		routes: make(map[string]types.HandlerFunc),
	}

	pc.InitHealthRoutes()

	return pc
}

func (p *PublicController) Route(
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
