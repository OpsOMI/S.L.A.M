package users

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"

	"github.com/OpsOMI/S.L.A.M/internal/shared/dto/users"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/response"
	"github.com/google/uuid"
)

func (s *module) Login(
	req *request.ClientRequest,
) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Password: ")
	bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", fmt.Errorf("failed to read password: %w", err)
	}
	fmt.Println()

	username = strings.TrimSpace(username)
	password := strings.TrimSpace(string(bytePassword))

	payload := users.LoginReq{
		Username: username,
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

	var data users.LoginResp
	if err := json.Unmarshal(dataBytes, &data); err != nil {
		return "", fmt.Errorf("failed to unmarshal login response data: %w", err)
	}

	return data.Token, nil
}

func (s *module) Register(
	req *request.ClientRequest,
) (uuid.UUID, error) {
	reader := bufio.NewReader(os.Stdin)

	// Get Nickname
	fmt.Print("Nickname: ")
	nickname, _ := reader.ReadString('\n')

	// Get Username
	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')

	// Get Password (hidden input)
	fmt.Print("Password: ")
	bytePassword, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to read password: %w", err)
	}
	fmt.Println()

	// Trim inputs
	nickname = strings.TrimSpace(nickname)
	username = strings.TrimSpace(username)
	password := strings.TrimSpace(string(bytePassword))

	// Create RegisterReq payload
	payload := users.RegisterReq{
		Nickname: nickname,
		Username: username,
		Password: password,
	}

	// Encode payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to encode payload: %w", err)
	}

	// Set payload into request and send
	req.Payload = payloadBytes
	if err := request.Send(s.conn, req); err != nil {
		return uuid.Nil, fmt.Errorf("failed to send register message: %w", err)
	}

	// Read response from server
	resp, err := response.Read(s.conn)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to read register response: %w", err)
	}

	baseResp := resp
	if baseResp.Code != "OK" {
		return uuid.Nil, fmt.Errorf("server error: %s", baseResp.Message)
	}

	dataBytes, err := json.Marshal(baseResp.Data)
	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to re-marshal login response data: %w", err)
	}

	var data users.RegisterResp
	if err := json.Unmarshal(dataBytes, &data); err != nil {
		return uuid.Nil, fmt.Errorf("failed to unmarshal login response data: %w", err)
	}

	return data.ID, nil
}
