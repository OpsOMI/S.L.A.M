package responder

import (
	"github.com/OpsOMI/S.L.A.M/internal/client/utils"
	"github.com/OpsOMI/S.L.A.M/internal/shared/dto/rooms"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/response"
)

func (r *Responder) HandleJoin(response response.BaseResponse) {
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

func (r *Responder) HandleCreateRoom(response response.BaseResponse) {
	var data rooms.CreateResp
	if err := utils.LoadData(response.Data, &data); err != nil {
		r.terminal.PrintError(err.Error())
		return
	}

	r.terminal.PrintNotification("Room Created Successfully, Code: " + data.Code)
}

func (r *Responder) HandleListRoom(response response.BaseResponse) {
	var data rooms.RoomsResp
	if err := utils.LoadData(response.Data, &data); err != nil {
		r.terminal.PrintError(err.Error())
		return
	}

	r.terminal.SetRooms(&data)
	r.terminal.PrintNotification("Your Rooms Listed!")
}
