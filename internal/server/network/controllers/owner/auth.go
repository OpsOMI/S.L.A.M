package owner

import (
	"context"
	"encoding/json"
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/server/domains/commons"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/response"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/utils"
	"github.com/OpsOMI/S.L.A.M/internal/shared/dto/users"
)

func (p *Controller) InitAuthRoutes() {
	p.routes["/register"] = p.HandleRegister
}

func (p *Controller) HandleRegister(
	conn net.Conn,
	args json.RawMessage,
	jwtToken *string,
) error {
	var req users.RegisterReq
	if err := utils.ParseJSON(args, &req); err != nil {
		return nil
	}

	ctx := context.Background()
	id, clientID, err := p.services.Users().CreateUser(ctx, req.Nickname, req.Username, req.Password, "user")
	if err != nil {
		return err
	}

	if p.cfg.Server.App.Mode == "prod" {
		if err := p.services.Clients().CreateClient(p.cfg, *clientID, req.Nickname); err != nil {
			return err
		}
	}

	return response.Response(commons.StatusOK, "User Created!", users.ToRegisterResponse(*id))
}
