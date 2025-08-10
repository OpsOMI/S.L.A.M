package response

import (
	"errors"
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/server/apperrors"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/response"
)

func Handle(conn net.Conn, err error, requestID string) error {
	if err == nil {
		return nil
	}

	// If error is already a BaseResponse (value or pointer), send it directly
	if respPtr, ok := err.(*response.BaseResponse); ok {
		respPtr.ReponseID = requestID
		return request.Send(conn, *respPtr)
	}

	// If error is an AppError, map its fields into BaseResponse
	var appErr *apperrors.AppError
	if errors.As(err, &appErr) {
		resp := response.BaseResponse{
			ReponseID: requestID,
			Message:   appErr.Message,
			Code:      appErr.Code,
		}

		// Hide internal details if error source is repository
		if appErr.Source == apperrors.SourceRepo {
			resp.Message = "Something Went Wrong"
			resp.Code = "Internal Server Error"
			resp.Data = nil
		}

		return request.Send(conn, resp)
	}

	// For unknown error types, send generic internal server error
	resp := response.BaseResponse{
		ReponseID: requestID,
		Message:   "Internal Server Error",
		Code:      "status.internal_server_error",
		Data:      nil,
	}
	return request.Send(conn, resp)
}

func Response(code, message string, data any) error {
	return &response.BaseResponse{
		Code:    code,
		Message: message,
		Data:    data,
		Errors:  nil,
	}
}
