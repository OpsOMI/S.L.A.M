package rooms

import (
	"time"

	"github.com/OpsOMI/S.L.A.M/internal/server/apperrors/domainerrors"
	"github.com/google/uuid"
)

type Room struct {
	ID        uuid.UUID
	OwnerID   uuid.UUID
	Code      string
	Password  string
	CreatedAt time.Time
}

type Rooms struct {
	Items      []Room
	TotalCount int64
}

// New constructs a new Room instance.
func New(
	ownerID uuid.UUID,
	code, password string,
) Room {
	return Room{
		OwnerID:  ownerID,
		Code:     code,
		Password: password,
	}
}

// ValidateCreate checks domain rules for client creation.
func (c *Room) ValidateCreate() error {
	if c.OwnerID == uuid.Nil {
		return domainerrors.BadRequest(ErrOwnerIDRequired)
	}
	if c.Code == "" {
		return domainerrors.BadRequest(ErrRoomCodeRequired)
	}

	return nil
}
