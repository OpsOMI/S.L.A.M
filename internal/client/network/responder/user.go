package responder

import (
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/response"
)

func (r *Responder) HandleMe(response response.BaseResponse) {
	r.terminal.PrintError(response.Message)

	// r.terminal.PrintNotification(response.Message)
}
