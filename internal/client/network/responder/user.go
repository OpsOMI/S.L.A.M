package responder

import (
	"fmt"

	"github.com/OpsOMI/S.L.A.M/internal/client/utils"
	"github.com/OpsOMI/S.L.A.M/internal/shared/dto/users"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/response"
)

func (r *Responder) HandleMe(response response.BaseResponse) {
	var data users.MeResp
	if err := utils.LoadData(response.Data, &data); err != nil {
		r.terminal.PrintError(err.Error())
		return
	}

	var maskedClientKey string
	if len(data.ClientKey) > 5 {
		maskedClientKey = data.ClientKey.String()[:5] + "..."
	}

	msg := fmt.Sprintf("ClientKey: %s, Nickname: %s, Username: %s", maskedClientKey, data.Nickname, data.Username)
	r.terminal.PrintNotification(msg, 3)
}

func (r *Responder) HandleOnline(response response.BaseResponse) {
	var data users.OnlineResp
	if err := utils.LoadData(response.Data, &data); err != nil {
		r.terminal.PrintError(err.Error())
		return
	}

	msg := fmt.Sprintf("Online User Count: %v", data.OnlineCount)
	r.terminal.PrintNotification(msg, 3)
}
