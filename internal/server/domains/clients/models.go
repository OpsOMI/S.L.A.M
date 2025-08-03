package clients

import (
	"time"

	"github.com/OpsOMI/S.L.A.M/internal/server/apperrors/domainerrors"
	"github.com/google/uuid"
)

type Client struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	ClientKey uuid.UUID
	RevokedAt *time.Time
	CreatedAt time.Time
}

type Clients struct {
	Items      []Client
	TotalCount int64
}

// New constructs a new Client instance.
func New(userID, clientKey uuid.UUID) Client {
	return Client{
		UserID:    userID,
		ClientKey: clientKey,
	}
}

// ValidateCreate checks domain rules for client creation.
func (c *Client) ValidateCreate() error {
	if c.UserID == uuid.Nil {
		return domainerrors.BadRequest(ErrUserIDRequired)
	}
	if c.ClientKey == uuid.Nil {
		return domainerrors.BadRequest(ErrClientKeyRequired)
	}

	return nil
}

// IsRevoked checks if the client is revoked.
func (c *Client) IsRevoked() bool {
	return c.RevokedAt != nil
}
