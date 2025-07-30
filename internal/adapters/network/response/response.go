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

// Success hem JSON oluşturur hem de conn'a yazar
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

// Error hem JSON oluşturur hem de conn'a yazar
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

// writeJSONResponse TCP bağlantısına JSON olarak yazmak için
func writeJSONResponse(conn net.Conn, jsonBytes []byte) error {
	_, err := conn.Write(append(jsonBytes, '\n'))
	return err
}
