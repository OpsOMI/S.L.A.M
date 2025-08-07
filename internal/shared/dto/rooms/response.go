package rooms

import "github.com/OpsOMI/S.L.A.M/internal/server/domains/rooms"

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
