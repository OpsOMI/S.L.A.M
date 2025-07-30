package public

import (
	"encoding/json"
	"fmt"
	"net"
)

func (p *PublicController) InitHealthRoutes() {
	p.routes["/ping"] = p.HandlePing
	p.routes["/info"] = p.HandleInfo
}

func (p *PublicController) HandlePing(conn net.Conn, args json.RawMessage) error {
	_, err := fmt.Fprintln(conn, `{"response":"pong"}`)
	return err
}

func (p *PublicController) HandleInfo(conn net.Conn, args json.RawMessage) error {
	_, err := fmt.Fprintln(conn, `{"info":"public controller info"}`)
	return err
}
