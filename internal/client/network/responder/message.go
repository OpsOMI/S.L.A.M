package responder

import (
	"github.com/OpsOMI/S.L.A.M/internal/client/utils"
	"github.com/OpsOMI/S.L.A.M/internal/shared/dto/message"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/response"
)

func (r *Responder) HandleSendMessage(response response.BaseResponse) {
	var data message.MessageResp
	if err := utils.LoadData(response.Data, &data); err != nil {
		r.terminal.PrintError(err.Error())
		return
	}

}

func (r *Responder) HandleIncomingMessages(response response.BaseResponse) {
	var data message.MessageResp
	if err := utils.LoadData(response.Data, &data); err != nil {
		r.terminal.PrintError(err.Error())
		return
	}

	if data.RoomCode == r.store.Room {
		r.terminal.AppendMessage(&data)
	}

	r.terminal.PrintNotification(response.Message + "geLDİ mSJ")
}
