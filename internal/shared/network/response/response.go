package response

import (
	"bufio"
	"encoding/json"
	"net"
)

type BaseResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Errors  any    `json:"errors,omitempty"`
	Data    any    `json:"details,omitempty"`
}

// Error implements the error interface for BaseResponse.
func (r *BaseResponse) Error() string {
	return r.Message
}

func Read(conn net.Conn) (*BaseResponse, error) {
	reader := bufio.NewReader(conn)
	line, err := reader.ReadBytes('\n')
	if err != nil {
		return nil, err
	}

	var resp BaseResponse
	err = json.Unmarshal(line, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
