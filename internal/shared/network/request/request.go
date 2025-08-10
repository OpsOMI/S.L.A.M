package request

import (
	"encoding/json"
	"net"
)

type ClientRequest struct {
	RequestID string          `json:"request_id"`
	JwtToken  string          `json:"jwt_token"`
	Scope     string          `json:"scope"`
	Command   string          `json:"command"`
	Payload   json.RawMessage `json:"payload"`
}

func Send(conn net.Conn, payload any) error {
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	_, err = conn.Write(append(b, '\n'))
	return err
}
