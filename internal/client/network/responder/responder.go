package responder

import (
	"github.com/OpsOMI/S.L.A.M/internal/client/network/commons"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/store"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/terminal"
	"github.com/OpsOMI/S.L.A.M/internal/client/utils"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/response"
)

type Responder struct {
	store    *store.SessionStore
	terminal *terminal.Terminal
}

func NewResponder(
	store *store.SessionStore,
	terminal *terminal.Terminal,
) Responder {
	return Responder{
		terminal: terminal,
		store:    store,
	}
}

func (r *Responder) Listen(responeseChan <-chan response.BaseResponse) {
	for response := range responeseChan {
		if err := utils.CheckBaseResponse(&response); err != nil {
			r.terminal.PrintError(err.Error())
			continue
		}

		switch response.ResponseID {
		case commons.RequestIDIncomingMessage:
			r.HandleIncomingMessages(response)
		case commons.RequestIDCleanRoom:
			r.HandleCleanRoom(response)
		case commons.RequestIDSendMessage:
			r.HandleSendMessage(response)
		case commons.RequestIDLogin:
			r.HandleLogin(response)
		case commons.RequestIDJoinRoom:
			r.HandleJoin(response)
		case commons.RequestIDCreateRoom:
			r.HandleCreateRoom(response)
		case commons.RequestIDListRoom:
			r.HandleListRoom(response)
		case commons.RequestIDRegister:
			r.HandleRegister(response)
		}
		r.terminal.Render()
	}
}
