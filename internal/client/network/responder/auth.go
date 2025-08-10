package responder

import (
	"github.com/OpsOMI/S.L.A.M/internal/client/utils"
	"github.com/OpsOMI/S.L.A.M/internal/shared/dto/users"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/response"
)

func (r *Responder) HandleLogin(response response.BaseResponse) {
	if err := utils.CheckBaseResponse(&response); err != nil {
		r.terminal.PrintError(err.Error())
		return
	}

	var data users.LoginResp
	if err := utils.LoadData(response.Data, &data); err != nil {
		r.terminal.PrintError(err.Error())
		return
	}

	r.store.SetToken(data.Token)
	r.store.ParseJWT()
	r.terminal.Render()

	r.terminal.PrintNotification("Login Successfull")
}
