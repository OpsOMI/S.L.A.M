package response

import (
	"encoding/json"
	"errors"
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/server/apperrors"
)

// BaseResponse defines the standard format for all server responses.
type BaseResponse struct {
	Message string      `json:"message"` // Message intended for the client
	Code    string      `json:"code"`    // forbidden, ok
	Errors  any         `json:"errors,omitempty"`
	Data    interface{} `json:"details,omitempty"` // Optional extra information
}

// Error implements the error interface for BaseResponse.
func (r *BaseResponse) Error() string {
	return r.Message
}

func Handle(conn net.Conn, err error) error {
	if err == nil {
		return nil
	}

	// If error is already a BaseResponse (value or pointer), send it directly
	if respPtr, ok := err.(*BaseResponse); ok {
		return WriteJson(conn, *respPtr)
	}

	// If error is an AppError, map its fields into BaseResponse
	var appErr *apperrors.AppError
	if errors.As(err, &appErr) {
		resp := BaseResponse{
			Message: appErr.Message,
			Code:    appErr.Code,
		}

		// Hide internal details if error source is repository
		if appErr.Source == apperrors.SourceRepo {
			resp.Message = "Something Went Wrong"
			resp.Code = "Internal Server Error"
			resp.Data = nil
		}

		return WriteJson(conn, resp)
	}

	// For unknown error types, send generic internal server error
	resp := BaseResponse{
		Message: "Internal Server Error",
		Code:    "status.internal_server_error",
		Data:    nil,
	}
	return WriteJson(conn, resp)
}

// WriteJson serializes the response and writes it to the connection.
func WriteJson(conn net.Conn, payload any) error {
	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	_, err = conn.Write(append(b, '\n'))
	return err
}

// Response creates and returns a BaseResponse representing a success response.
func Response(code, message string, data any) error {
	return &BaseResponse{
		Code:    code,
		Message: message,
		Data:    data,
		Errors:  nil,
	}
}
