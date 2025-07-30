package public

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/response"
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
) {
	handler, ok := p.routes[cmd]
	if !ok {
		_ = response.Error(conn, fmt.Sprintf("unknown command: %s", cmd))
	}

	// Here, we handle server-side errors such as I/O or JSON parsing issues.
	// Errors related to client requests should be handled and communicated by the handler itself.
	// If the handler returns an error, it means something went wrong on the server side, which we log for debugging.
	err := handler(conn, args, nil)
	if err != nil {
		p.logger.Error("Handler error: " + err.Error())
	}
}
