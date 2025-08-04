package private

import (
	"encoding/json"
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/server/domains/commons"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/response"
)

func (p *Controller) InitUserRoutes() {
	p.routes["/me"] = p.HandleMe
}

func (p *Controller) HandleMe(
	conn net.Conn,
	args json.RawMessage,
	jwtToken *string,
) error {
	userInfo := p.store.ParseToken(jwtToken)

	return response.Response(commons.StatusOK, "Me Command Worked!", userInfo)
}
