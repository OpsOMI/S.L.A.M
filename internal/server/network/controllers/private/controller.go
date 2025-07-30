package private

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"
	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/response"
	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/tokenstore"
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
) {
	handler, ok := p.routes[cmd]
	if !ok {
		_ = response.Error(conn, fmt.Sprintf("unknown command: %s", cmd))
	}

	finalHandler := handler
	for i := len(p.middlewares) - 1; i >= 0; i-- {
		finalHandler = p.middlewares[i](finalHandler)
	}

	// Here, we handle server-side errors such as I/O or JSON parsing issues.
	// Errors related to client requests should be handled and communicated by the handler itself.
	// If the handler returns an error, it means something went wrong on the server side, which we log for debugging.
	err := finalHandler(conn, args, &jwtToken)
	if err != nil {
		p.logger.Error("Handler error: " + err.Error())
	}
}
