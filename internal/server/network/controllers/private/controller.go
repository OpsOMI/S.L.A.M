package private

import (
	"encoding/json"
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/response"
	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/tokenstore"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/commons"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/middlewares"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/types"
)

type PrivateController struct {
	logger      logger.ILogger
	routes      map[string]types.HandlerFunc
	tokenstore  tokenstore.ITokenStore
	middlewares []types.Middleware
}

func NewController(
	logger logger.ILogger,
	tokenstore tokenstore.ITokenStore,
) *PrivateController {
	pc := &PrivateController{
		tokenstore:  tokenstore,
		logger:      logger,
		routes:      make(map[string]types.HandlerFunc),
		middlewares: make([]types.Middleware, 0),
	}

	pc.Use(middlewares.JWTAuthMiddleware(tokenstore))
	pc.InitUserRoutes()

	return pc
}

func (p *PrivateController) Use(mw types.Middleware) {
	p.middlewares = append(p.middlewares, mw)
}

func (p *PrivateController) Route(
	conn net.Conn,
	jwtToken, cmd string,
	args json.RawMessage,
) error {
	handler, ok := p.routes[cmd]
	if !ok {
		return response.Response(commons.StatusBadRequest, "Unknown Command", cmd)
	}

	finalHandler := handler
	for i := len(p.middlewares) - 1; i >= 0; i-- {
		finalHandler = p.middlewares[i](finalHandler)
	}

	return finalHandler(conn, args, &jwtToken)
}
