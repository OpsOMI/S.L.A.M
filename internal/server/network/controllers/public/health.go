package public

import (
	"encoding/json"
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/response"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/commons"
)

func (p *Controller) InitHealthRoutes() {
	p.routes["/ping"] = p.HandlePing
	p.routes["/"] = p.HandleRoot
}

func (p *Controller) HandleRoot(
	conn net.Conn,
	args json.RawMessage,
	jwtToken *string,
) error {
	return response.Response(commons.StatusOK, "Welcome To The Public Controller", nil)
}

func (p *Controller) HandlePing(
	conn net.Conn,
	args json.RawMessage,
	jwtToken *string,
) error {
	return response.Response(commons.StatusOK, "PONG", nil)
}
