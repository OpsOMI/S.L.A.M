package rooms

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"

	"github.com/OpsOMI/S.L.A.M/internal/shared/dto/rooms"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/response"
)

func (s *module) MyRooms(
	req *request.ClientRequest,
	page, limit int32,
) (*rooms.RoomsResp, error) {
	payload := rooms.MyRoomReq{
		Page:  page,
		Limit: limit,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to encode payload: %w", err)
	}

	req.Payload = payloadBytes

	if err := request.Send(s.conn, req); err != nil {
		return nil, fmt.Errorf("failed to send login message: %w", err)
	}

	resp, err := response.Read(s.conn)
	if err != nil {
		return nil, fmt.Errorf("failed to read login response: %w", err)
	}

	baseResp := resp
	if baseResp.Code != "OK" {
		return nil, fmt.Errorf("server error: %s", baseResp.Message)
	}

	dataBytes, err := json.Marshal(baseResp.Data)
	if err != nil {
		return nil, fmt.Errorf("failed to re-marshal login response data: %w", err)
	}

	var data rooms.RoomsResp
	if err := json.Unmarshal(dataBytes, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal login response data: %w", err)
	}

	return &data, nil
}

func (s *module) Create(
	req *request.ClientRequest,
	isSecure bool,
) (string, error) {
	var password string
	if isSecure {

		fmt.Print("Password: ")
		bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			return "", fmt.Errorf("failed to read password: %w", err)
		}
		fmt.Println()

		fmt.Print("Confirm Password: ")
		byteConfirmPassword, err := term.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			return "", fmt.Errorf("failed to read password: %w", err)
		}
		fmt.Println()

		confirmPassword := strings.TrimSpace(string(byteConfirmPassword))
		password = strings.TrimSpace(string(bytePassword))

		if password != confirmPassword {
			return "", fmt.Errorf("passwords do not match")
		}
	}

	payload := rooms.CreateReq{
		Password: password,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to encode payload: %w", err)
	}

	req.Payload = payloadBytes

	if err := request.Send(s.conn, req); err != nil {
		return "", fmt.Errorf("failed to send login message: %w", err)
	}

	resp, err := response.Read(s.conn)
	if err != nil {
		return "", fmt.Errorf("failed to read login response: %w", err)
	}

	baseResp := resp
	if baseResp.Code != "OK" {
		return "", fmt.Errorf("server error: %s", baseResp.Message)
	}

	dataBytes, err := json.Marshal(baseResp.Data)
	if err != nil {
		return "", fmt.Errorf("failed to re-marshal login response data: %w", err)
	}

	var data rooms.CreateResp
	if err := json.Unmarshal(dataBytes, &data); err != nil {
		return "", fmt.Errorf("failed to unmarshal login response data: %w", err)
	}

	return data.Code, nil
}
