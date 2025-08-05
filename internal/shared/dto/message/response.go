package message

import "github.com/OpsOMI/S.L.A.M/internal/server/domains/messages"

type MessageResp struct {
	RoomCode       string `json:"roomCode"`
	SenderNickname string `json:"senderNickname"`
	Content        string `json:"content"`
}

type MessagesReps struct {
	Items      []MessageResp `json:"items"`
	TotalCount int64         `json:"totalCount"`
}

func OneMessage(
	domainModel *messages.RoomMessage,
) *MessageResp {
	if domainModel == nil {
		return nil
	}
	return &MessageResp{
		SenderNickname: domainModel.SenderNickname,
		Content:        domainModel.ContentEnc,
	}
}

func ManyMessage(
	domainModel *messages.RoomMessages,
) MessagesReps {
	result := make([]MessageResp, 0, len(domainModel.Items))

	for _, msg := range domainModel.Items {
		result = append(result, *OneMessage(&msg))
	}

	return MessagesReps{
		Items:      result,
		TotalCount: domainModel.TotalCount,
	}
}
