package response

import (
	"encoding/json"
	"net"
)

type Response struct {
	Status  string `json:"status"`            // "success" or "error"
	Message string `json:"message,omitempty"` // message
	Data    any    `json:"data,omitempty"`    // Data
}

func Success(conn net.Conn, data any) error {
	resp := Response{
		Status: "success",
		Data:   data,
	}

	jsonBytes, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	return writeJSONResponse(conn, jsonBytes)
}

func Error(conn net.Conn, message string) error {
	resp := Response{
		Status:  "error",
		Message: message,
	}

	jsonBytes, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	return writeJSONResponse(conn, jsonBytes)
}

// Forbidden sends a standard forbidden error message to the client
func Forbidden(conn net.Conn) error {
	resp := Response{
		Status:  "error",
		Message: "forbidden",
	}

	jsonBytes, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	return writeJSONResponse(conn, jsonBytes)
}

func writeJSONResponse(conn net.Conn, jsonBytes []byte) error {
	_, err := conn.Write(append(jsonBytes, '\n'))
	return err
}
