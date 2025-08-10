package responder

import (
	"github.com/OpsOMI/S.L.A.M/internal/client/network/commons"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/store"
	"github.com/OpsOMI/S.L.A.M/internal/client/network/terminal"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/response"
)

type Responder struct {
	terminal *terminal.Terminal
	store    *store.SessionStore
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
		switch response.ResponseID {
		case commons.RequestIDLogin:
			r.HandleLogin(response)
		case commons.RequestIDJoin:
			r.HandleJoin(response)
		}
	}
}
