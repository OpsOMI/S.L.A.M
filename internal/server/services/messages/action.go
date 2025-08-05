package messages

import (
	"context"

	"github.com/OpsOMI/S.L.A.M/internal/server/apperrors/repoerrors"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/messages"
	"github.com/google/uuid"
)

func (s *service) GetMessagesByRoomCode(
	ctx context.Context,
	roomCode string,
) (*messages.RoomMessages, error) {
	return s.repositories.Messages().GetMessagesByRoomCode(ctx, roomCode)
}

func (s *service) CreateMessage(
	ctx context.Context,
	senderID, receiverID, roomCode string,
	conent string,
) error {
	sender, err := s.users.GetByID(ctx, senderID)
	if err != nil {
		return err
	}

	var rcid, rid *uuid.UUID
	if receiverID != "" {
		receiver, err := s.users.GetByID(ctx, receiverID)
		if err != nil {
			return err
		}
		rcid = &receiver.ID
	}
	if roomCode != "" {
		rooms, err := s.rooms.GetByCode(ctx, roomCode)
		if err != nil {
			return err
		}
		rid = &rooms.ID
	}

	if err := s.repositories.Messages().CreateMessage(ctx, sender.ID, rcid, rid, conent); err != nil {
		return repoerrors.Internal(messages.ErrCreateFailed, err)
	}

	return nil
}
