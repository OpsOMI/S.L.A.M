package types

import (
	"encoding/json"
	"net"
)

type HandlerFunc func(conn net.Conn, args json.RawMessage, jwtToken *string) error
type Middleware func(next HandlerFunc) HandlerFunc
