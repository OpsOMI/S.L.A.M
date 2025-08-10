package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/OpsOMI/S.L.A.M/internal/client/apperrors"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/response"
	"golang.org/x/term"
)

func SendRequest(
	conn net.Conn,
	req *request.ClientRequest,
	payload any,
) (*response.BaseResponse, error) {
	if payload != nil {
		payloadBytes, err := marshal(payload)
		if err != nil {
			return nil, err
		}
		req.Payload = payloadBytes
	}

	if err := request.Send(conn, req); err != nil {
		return nil, apperrors.NewError("failed to send message: " + err.Error())
	}

	// resp, err := response.Read(conn)
	// if err != nil {
	// 	return nil, apperrors.NewError("failed to read server response: " + err.Error())
	// }

	return nil, nil
}

func ResponseRead(conn net.Conn) (*response.BaseResponse, error) {
	resp, err := response.Read(conn)
	if err != nil {
		return nil, apperrors.NewError("failed to read server response: " + err.Error())
	}

	return resp, nil
}

func LoadData[T any](data any, target *T) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return apperrors.NewError("failed to re-marshal login response data: " + err.Error())
	}

	if err := json.Unmarshal(dataBytes, target); err != nil {
		return apperrors.NewError("failed to decode payload: " + err.Error())
	}

	return nil
}

// If the code OK returns nil.
func CheckBaseResponse(resp *response.BaseResponse) error {
	// if resp == nil {
	// 	return apperrors.NewError("empty server response")
	// }

	switch resp.Code {
	case "OK":
		return nil
	case "BadRequest":
		return apperrors.NewError("bad request: " + resp.Message)
	case "Unauthorized":
		return apperrors.NewError("unauthorized: " + resp.Message)
	default:
		return apperrors.NewError("server error: " + resp.Message)
	}
}

// Reads
func ReadPassword(label string) (string, error) {
	fmt.Printf("%s: ", label)
	bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", apperrors.NewError("failed to read password: " + err.Error())
	}

	return strings.TrimSpace(string(bytePassword)), nil
}

func Read(reader *bufio.Reader, label string) (string, error) {
	fmt.Printf("%s: ", label)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", apperrors.NewError("failed to read input: " + err.Error())
	}

	return strings.TrimSpace(input), nil
}

func marshal(payload any) ([]byte, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, apperrors.NewError("failed to encode payload: " + err.Error())
	}
	return payloadBytes, nil
}
