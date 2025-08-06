package messages

import (
	"github.com/OpsOMI/S.L.A.M/internal/adapters/postgres/sqlc/pgqueries"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/utils"
	"github.com/google/uuid"
)

type IMessagesMapper interface {
	One(
		dbModel *pgqueries.Message,
	) *Message

	OneRoomMessage(
		dbModel *pgqueries.GetMessagesByRoomCodeRow,
	) *RoomMessage

	Many(
		dbModels []pgqueries.Message,
		count int64,
	) *Messages

	ManyRoomMessages(
		dbModels []pgqueries.GetMessagesByRoomCodeRow,
		count int64,
	) *RoomMessages

	CreateParams(
		senderID, roomID uuid.UUID,
		content string,
	) pgqueries.CreateMessageParams
}

type mapper struct {
	utils utils.IUtilMapper
}

// NewMapper creates a new instance of messages mapper.
func NewMapper(utils utils.IUtilMapper) *mapper {
	return &mapper{
		utils: utils,
	}
}

// One maps a pgqueries.Message to a domain Message.
func (m *mapper) One(
	dbModel *pgqueries.Message,
) *Message {
	if dbModel == nil {
		return nil
	}

	return &Message{
		ID:         dbModel.ID,
		SenderID:   dbModel.SenderID,
		RoomID:     dbModel.RoomID,
		ContentEnc: dbModel.ContentEnc,
		CreatedAt:  dbModel.CreatedAt.Time,
	}
}

func (m *mapper) Many(
	dbModels []pgqueries.Message,
	count int64,
) *Messages {
	var appModels []Message
	for _, dbModel := range dbModels {
		appModels = append(appModels, *m.One(&dbModel))
	}

	return &Messages{
		Items:      appModels,
		TotalCount: count,
	}
}

func (m *mapper) OneRoomMessage(
	dbModel *pgqueries.GetMessagesByRoomCodeRow,
) *RoomMessage {
	if dbModel == nil {
		return nil
	}

	return &RoomMessage{
		SenderNickname: dbModel.SenderNickname,
		ContentEnc:     dbModel.ContentEnc,
	}
}

func (m *mapper) ManyRoomMessages(
	dbModels []pgqueries.GetMessagesByRoomCodeRow,
	count int64,
) *RoomMessages {
	var appModels []RoomMessage
	for _, dbModel := range dbModels {
		appModels = append(appModels, *m.OneRoomMessage(&dbModel))
	}

	return &RoomMessages{
		Items:      appModels,
		TotalCount: count,
	}
}

func (m *mapper) CreateParams(
	senderID, roomID uuid.UUID,
	content string,
) pgqueries.CreateMessageParams {
	return pgqueries.CreateMessageParams{
		SenderID: senderID,
		RoomID:   roomID,
		Content:  content,
	}
}
