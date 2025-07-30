package public

import (
	"encoding/json"
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/response"
)

func (p *PublicController) InitHealthRoutes() {
	p.routes["/ping"] = p.HandlePing
	p.routes["/info"] = p.HandleRoot
}

func (p *PublicController) HandleRoot(
	conn net.Conn,
	args json.RawMessage,
	jwtToken *string,
) error {
	return response.Success(conn, "Welcome To The Public Controller")
}

func (p *PublicController) HandlePing(
	conn net.Conn,
	args json.RawMessage,
	jwtToken *string,
) error {
	return response.Success(conn, "PONG")
}
