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

func (s *service) Login() error {
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
		return fmt.Errorf("failed to encode payload: %w", err)
	}

	msg := request.ClientMessage{
		Command: "/auth/login",
		Scope:   "public",
		Payload: payloadBytes,
	}

	if err := request.Send(s.conn, msg); err != nil {
		return fmt.Errorf("failed to send login message: %w", err)
	}

	if err := request.Send(s.conn, msg); err != nil {
		return fmt.Errorf("failed to send login message: %w", err)
	}

	resp, err := response.Read(s.conn)
	if err != nil {
		return fmt.Errorf("failed to read login response: %w", err)
	}

	token, ok := resp.Data.(string)
	if !ok {
		return fmt.Errorf("token is not a string")
	}

	// Store token somewhere (e.g., global, file, memory)
	fmt.Println("Login successful.")
	fmt.Println("JWT Token:", token)

	return nil
}
