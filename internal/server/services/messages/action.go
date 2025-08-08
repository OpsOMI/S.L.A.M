package messages

import (
	"context"

	"github.com/OpsOMI/S.L.A.M/internal/server/apperrors/repoerrors"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/messages"
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
		return repoerrors.Internal(messages.ErrCreateFailed, err)
	}

	return nil
}

func (s *service) DeleteMessages(
	ctx context.Context,
) error {
	return s.repositories.Messages().DeleteMessages(ctx)
}

func (s *service) DeleteMessageByRoomCode(
	ctx context.Context,
	roomCode string,
) error {
	return s.repositories.Messages().DeleteMessagesByRoomCode(ctx, roomCode)
}
