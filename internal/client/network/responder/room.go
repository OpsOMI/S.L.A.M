package responder

import (
	"github.com/OpsOMI/S.L.A.M/internal/client/utils"
	"github.com/OpsOMI/S.L.A.M/internal/shared/dto/rooms"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/response"
)

func (r *Responder) HandleJoin(response response.BaseResponse) {
	if err := utils.CheckBaseResponse(&response); err != nil {
		r.terminal.PrintError(err.Error())
		return
	}

	var data rooms.JoinResp
	if err := utils.LoadData(response.Data, &data); err != nil {
		r.terminal.PrintError(err.Error())
		return
	}

	r.store.SetRoom(data.RoomCode)
	r.terminal.SetCurrentRoom(data.RoomCode)
	r.terminal.SetMessages(&data.Messages)

	r.terminal.PrintNotification("You joined room " + data.RoomCode)
}
