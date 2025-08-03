package users

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/request"
	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/response"
)

func (s *module) Login(req *request.ClientRequest) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Username: ")
	username, _ := reader.ReadString('\n')

	fmt.Print("Password: ")
	password, _ := reader.ReadString('\n')

	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	payload := map[string]string{
		"username": username,
		"password": password,
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

	token, ok := baseResp.Data.(string)
	if !ok {
		tokenBytes, err := json.Marshal(baseResp.Data)
		if err != nil {
			return "", fmt.Errorf("failed to marshal token data: %w", err)
		}
		token = string(tokenBytes)
	}

	return token, nil
}
