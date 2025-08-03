package public

import (
	"encoding/json"
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/response"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/commons"
)

func (p *Controller) InitAuthRoutes() {
	p.routes["/auth/login"] = p.HandleLogin
}

func (p *Controller) HandleLogin(
	conn net.Conn,
	args json.RawMessage,
	jwtToken *string,
) error {
	return response.Response(commons.StatusOK, "PONG", nil)
}
