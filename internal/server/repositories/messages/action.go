package messages

import (
	"context"

	"github.com/OpsOMI/S.L.A.M/internal/server/apperrors/repoerrors"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/messages"
	"github.com/google/uuid"
)

func (r *repository) GetMessagesByRoomCode(
	ctx context.Context,
	roomCode string,
) (*messages.RoomMessages, error) {
	dbModels, err := r.queries.GetMessagesByRoomCode(ctx, roomCode)
	if err != nil {
		return nil, repoerrors.Internal(messages.ErrFetchFailed, err)
	}

	count, err := r.queries.CountMessagesByRoomCode(ctx, roomCode)
	if err != nil {
		return nil, repoerrors.Internal(messages.ErrCountFailed, err)
	}

	return r.mappers.Messages().ManyRoomMessages(dbModels, count), nil
}

func (r *repository) CreateMessage(
	ctx context.Context,
	senderID uuid.UUID,
	receiverID, roomID *uuid.UUID,
	content string,
) error {
	params := r.mappers.Messages().CreateParams(senderID, receiverID, roomID, content)

	if err := r.queries.CreateMessage(ctx, params); err != nil {
		return repoerrors.Internal(messages.ErrCreateFailed, err)
	}

	return nil
}
