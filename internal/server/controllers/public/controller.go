package public

import (
	"encoding/json"
	"fmt"
	"net"
)

type HandlerFunc func(conn net.Conn, args json.RawMessage) error

type PublicController struct {
	routes map[string]HandlerFunc
}

func NewController() *PublicController {
	pc := &PublicController{
		routes: make(map[string]HandlerFunc),
	}

	pc.InitHealthRoutes()

	return pc
}

func (p *PublicController) Route(
	conn net.Conn,
	cmd string,
	args json.RawMessage,
) error {
	handler, ok := p.routes[cmd]
	if !ok {
		return fmt.Errorf("public controller: unknown command %s", cmd)
	}

	return handler(conn, args)
}
