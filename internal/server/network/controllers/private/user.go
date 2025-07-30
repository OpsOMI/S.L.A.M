package private

import (
	"encoding/json"
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/response"
)

func (p *PrivateController) InitUserRoutes() {
	p.routes["/me"] = p.HandleMe
}

func (p *PrivateController) HandleMe(
	conn net.Conn,
	args json.RawMessage,
	jwtToken *string,
) error {
	userInfo := p.tokenstore.ParseToken(jwtToken)

	return response.Success(conn, userInfo)
}
