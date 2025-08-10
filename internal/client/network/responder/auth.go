package responder

import (
	"github.com/OpsOMI/S.L.A.M/internal/client/utils"
	"github.com/OpsOMI/S.L.A.M/internal/shared/dto/users"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/response"
)

func (r *Responder) HandleLogin(response response.BaseResponse) {
	var data users.LoginResp
	if err := utils.LoadData(response.Data, &data); err != nil {
		r.terminal.PrintError(err.Error())
		return
	}

	r.store.SetToken(data.Token)
	r.store.ParseJWT()
	r.terminal.SetPromptLabel("->", r.store.Nickname)

	r.terminal.PrintNotification("Login Successfull")
}

func (r *Responder) HandleRegister(response response.BaseResponse) {
	var data users.RegisterResp
	if err := utils.LoadData(response.Data, &data); err != nil {
		r.terminal.PrintError(err.Error())
		return
	}

	r.terminal.PrintNotification("New User Created: " + data.ID.String())
}
