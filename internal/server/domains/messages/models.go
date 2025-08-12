package messages

import (
	"time"

	"github.com/OpsOMI/S.L.A.M/internal/server/apperrors/domainerrors"
	"github.com/google/uuid"
)

type Message struct {
	ID         uuid.UUID
	SenderID   uuid.UUID
	RoomID     uuid.UUID
	ContentEnc string
	CreatedAt  time.Time
}

type Messages struct {
	Items      []Message
	TotalCount int64
}

type RoomMessage struct {
	SenderNickname string
	ContentEnc     string
}

type RoomMessages struct {
	Items      []*RoomMessage
	TotalCount int64
}

func New(
	senderID, roomID uuid.UUID,
	contentEnc string,
) Message {
	return Message{
		SenderID:   senderID,
		RoomID:     roomID,
		ContentEnc: contentEnc,
	}
}

func (m *Message) ValidateCreate() error {
	if m.SenderID == uuid.Nil {
		return domainerrors.BadRequest(ErrSenderIDRequired)
	}
	if m.RoomID == uuid.Nil {
		return domainerrors.BadRequest(ErrRoomIDRequired)
	}

	return nil
}
