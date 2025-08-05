package messages

import (
	"context"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/postgres/sqlc/pgqueries"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/messages"
	"github.com/OpsOMI/S.L.A.M/pkg/txmanagerpkg"
	"github.com/google/uuid"
)

type IMessagesRepository interface {
	GetMessagesByRoomCode(
		ctx context.Context,
		roomCode string,
	) (*messages.RoomMessages, error)

	CreateMessage(
		ctx context.Context,
		senderID uuid.UUID,
		receiverID, roomID *uuid.UUID,
		content string,
	) error
}

type repository struct {
	queries   *pgqueries.Queries
	txManager txmanagerpkg.ITXManager
	mappers   domains.IMapper
}

func NewRepository(
	queries *pgqueries.Queries,
	mappers domains.IMapper,
	txManager txmanagerpkg.ITXManager,
) IMessagesRepository {
	return &repository{
		txManager: txManager,
		queries:   queries,
		mappers:   mappers,
	}
}
