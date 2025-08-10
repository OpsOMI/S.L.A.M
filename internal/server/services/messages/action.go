package messages

import (
	"context"
	"strings"

	"github.com/OpsOMI/S.L.A.M/internal/server/apperrors/serviceerrors"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/messages"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/rooms"
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
	senderID, roomCode string,
	content string,
) error {
	sender, err := s.users.GetByID(ctx, senderID)
	if err != nil {
		return err
	}

	rooms, err := s.rooms.GetByCode(ctx, roomCode)
	if err != nil {
		return err
	}

	if err := s.repositories.Messages().CreateMessage(ctx, sender.ID, rooms.ID, content); err != nil {
		return serviceerrors.Internal(messages.ErrCreateFailed, err)
	}

	return nil
}

func (s *service) DeleteMessages(
	ctx context.Context,
) error {
	return s.repositories.Messages().DeleteMessages(ctx)
}

func (s *service) DeleteMessageInRoom(
	ctx context.Context,
	ownerID uuid.UUID,
	roomCode string,
) error {
	if _, err := s.rooms.GetByCodeAndOwnerID(ctx, ownerID.String(), roomCode); err != nil {
		if strings.Contains(err.Error(), "not_found") {
			return serviceerrors.Forbidden(rooms.ErrNotYourRoom)
		}
		return err
	}

	return s.repositories.Messages().DeleteMessagesByRoomCode(ctx, roomCode)
}

func (s *service) DeleteMessageByRoomCode(
	ctx context.Context,
	roomCode string,
) error {
	return s.repositories.Messages().DeleteMessagesByRoomCode(ctx, roomCode)
}
