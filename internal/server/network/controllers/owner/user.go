package owner

import (
	"encoding/json"
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/server/domains/commons"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/response"
	"github.com/OpsOMI/S.L.A.M/internal/shared/dto/users"
)

func (p *Controller) InitUserRoutes() {
	p.routes["/online"] = p.HandleOnline
}

func (p *Controller) HandleOnline(
	conn net.Conn,
	args json.RawMessage,
	jwtToken *string,
) error {
	onlineCounts := p.connections.CountOnlineConnections()

	return response.Response(commons.StatusOK, "Online Connections", users.ToOnlineResp(onlineCounts))
}
