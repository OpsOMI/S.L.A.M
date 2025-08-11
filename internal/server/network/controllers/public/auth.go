package public

import (
	"context"
	"encoding/json"
	"net"
	"time"

	"github.com/OpsOMI/S.L.A.M/internal/server/domains/commons"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/response"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/utils"
	"github.com/OpsOMI/S.L.A.M/internal/shared/dto/client"
	"github.com/OpsOMI/S.L.A.M/internal/shared/dto/users"
)

func (p *Controller) InitAuthRoutes() {
	p.routes["/login"] = p.HandleLogin
}

func (p *Controller) HandleLogin(
	conn net.Conn,
	args json.RawMessage,
	jwtToken *string,
) error {
	var req users.LoginReq
	if err := utils.ParseJSON(args, &req); err != nil {
		return err
	}

	ctx := context.Background()
	user, err := p.services.Users().Login(ctx, req.ClientKey, req.Username, req.Password, p.cfg.Server.App.Mode)
	if err != nil {
		return err
	}

	client, err := p.services.Clients().GetByClientKey(ctx, req.ClientKey)
	if err != nil {
		return err
	}

	jwt, err := p.store.GenerateToken(client.ClientKey, user.ID, user.Username, user.Nickname, user.Role, 24*time.Hour)
	if err != nil {
		return err
	}

	p.connecions.Register(client.ClientKey, conn)

	return response.Response(commons.StatusOK, "Login successful", users.ToLoginResponse(jwt))
}

func (p *Controller) HandleClient(
	conn net.Conn,
	args json.RawMessage,
	jwtToken *string,
) error {
	var req client.ClientReq
	if err := utils.ParseJSON(args, &req); err != nil {
		return err
	}

	ctx := context.Background()
	isExists, err := p.services.Clients().IsExistByClientKey(ctx, req.ClientKey)
	if err != nil {
		return err
	}

	return response.Response(commons.StatusOK, "Login successful", client.ToClientResp(*isExists))
}
