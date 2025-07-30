package request

import "encoding/json"

type ClientMessage struct {
	JwtToken string          `json:"jwt_token"` // JWT token for authentication and authorization
	Scope    string          `json:"scope"`     // User scope or role, e.g., "public", "private", "owner"
	Command  string          `json:"command"`   // Command to execute, e.g., "/join", "/message"
	Payload  json.RawMessage `json:"payload"`   // Command-specific data in JSON format
}
