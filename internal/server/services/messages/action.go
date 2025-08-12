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
	roomCode, secret string,
) (*messages.RoomMessages, error) {
	models, err := s.repositories.Messages().GetMessagesByRoomCode(ctx, roomCode)
	if err != nil {
		return nil, err
	}

	for _, ms := range models.Items {
		content, err := s.packages.Hasher().DecryptMessage(ms.ContentEnc, []byte(secret))
		if err != nil {
			return nil, serviceerrors.Internal(messages.ErrMessageDecryptFailed, err)
		}
		ms.ContentEnc = content
	}

	return models, nil
}

func (s *service) CreateMessage(
	ctx context.Context,
	senderID, roomCode, content, secret string,
) error {
	sender, err := s.users.GetByID(ctx, senderID)
	if err != nil {
		return err
	}

	rooms, err := s.rooms.GetByCode(ctx, roomCode)
	if err != nil {
		return err
	}

	hashedContent, err := s.packages.Hasher().EncryptMessage(content, []byte(secret))
	if err != nil {
		return serviceerrors.Internal(messages.ErrMessageEncryptFailed, err)
	}

	if err := s.repositories.Messages().CreateMessage(ctx, sender.ID, rooms.ID, hashedContent); err != nil {
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
