package rooms

import (
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/messages"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/rooms"
	"github.com/OpsOMI/S.L.A.M/internal/shared/dto/message"
)

type CreateResp struct {
	Code string `json:"code"`
}

func OneCreate(
	code string,
) CreateResp {
	return CreateResp{
		Code: code,
	}
}

type RoomResp struct {
	Code     string `json:"code"`
	IsLocked bool   `json:"isLocked"`
}

type RoomsResp struct {
	Items      []RoomResp `json:"items"`
	TotalCount int64      `json:"totalCount"`
}

func OneRoom(
	room *rooms.Room,
) RoomResp {
	return RoomResp{
		Code:     room.Code,
		IsLocked: room.Password != "",
	}
}

func ManyRoom(
	rooms *rooms.Rooms,
) RoomsResp {
	var items []RoomResp
	for _, model := range rooms.Items {
		items = append(items, OneRoom(&model))
	}

	return RoomsResp{
		Items:      items,
		TotalCount: rooms.TotalCount,
	}
}

type JoinResp struct {
	RoomCode string               `json:"roomCode"`
	Messages message.MessagesReps `json:"messages"`
}

func OneJoin(
	RoomCode string,
	domainModel *messages.RoomMessages,
) JoinResp {
	return JoinResp{
		RoomCode: RoomCode,
		Messages: message.ManyMessage(domainModel),
	}
}
