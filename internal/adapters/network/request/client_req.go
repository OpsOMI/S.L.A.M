package request

import (
	"encoding/json"

	"github.com/OpsOMI/S.L.A.M/internal/server/apperrors/serviceerrors"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/commons"
)

type ClientMessage struct {
	JwtToken string          `json:"jwt_token"` // JWT token for authentication and authorization
	Scope    string          `json:"scope"`     // User scope or role, e.g., "public", "private", "owner"
	Command  string          `json:"command"`   // Command to execute, e.g., "/join", "/message"
	Payload  json.RawMessage `json:"payload"`   // Command-specific data in JSON format
}

func ParseJSON[T any](data json.RawMessage, target *T) error {
	if err := json.Unmarshal(data, target); err != nil {
		return serviceerrors.BadRequest(commons.ErrParseFailed)
	}
	return nil
}
