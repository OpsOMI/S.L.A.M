package private

import (
	"encoding/json"
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/logger"

	"github.com/OpsOMI/S.L.A.M/internal/server/config"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/commons"
	"github.com/OpsOMI/S.L.A.M/internal/server/infrastructure/connection"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/middlewares"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/response"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/store"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/types"
	"github.com/OpsOMI/S.L.A.M/internal/server/services"
)

type Controller struct {
	cfg         *config.Configs
	logger      logger.ILogger
	routes      map[string]types.HandlerFunc
	store       store.IJwtManager
	middlewares []types.Middleware
	services    services.IServices
	connections *connection.ConnectionManager
}

func NewController(
	cfg *config.Configs,
	logger logger.ILogger,
	store store.IJwtManager,
	services services.IServices,
	connections *connection.ConnectionManager,
) *Controller {
	pc := &Controller{
		cfg:         cfg,
		store:       store,
		logger:      logger,
		routes:      make(map[string]types.HandlerFunc),
		middlewares: make([]types.Middleware, 0),
		services:    services,
		connections: connections,
	}

	pc.Use(middlewares.JWTAuthMiddleware(store))
	pc.InitUserRoutes()
	pc.InitRoomRoutes()

	return pc
}

func (p *Controller) Use(mw types.Middleware) {
	p.middlewares = append(p.middlewares, mw)
}

func (p *Controller) Route(
	conn net.Conn,
	jwtToken, cmd string,
	args json.RawMessage,
) error {
	handler, ok := p.routes[cmd]
	if !ok {
		return response.Response(commons.StatusBadRequest, "Unknown Command: "+cmd, nil)
	}

	finalHandler := handler
	for i := len(p.middlewares) - 1; i >= 0; i-- {
		finalHandler = p.middlewares[i](finalHandler)
	}

	return finalHandler(conn, args, &jwtToken)
}
