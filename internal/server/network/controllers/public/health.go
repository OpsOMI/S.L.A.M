package public

import (
	"encoding/json"
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/server/domains/commons"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/response"
)

func (p *Controller) InitHealthRoutes() {
	p.routes["/"] = p.HandleRoot
}

func (p *Controller) HandleRoot(
	conn net.Conn,
	args json.RawMessage,
	jwtToken *string,
) error {
	return response.Response(commons.StatusOK, "Welcome To The Public Controller", nil)
}
