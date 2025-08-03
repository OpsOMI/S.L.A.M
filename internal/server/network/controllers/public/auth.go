package public

import (
	"context"
	"encoding/json"
	"net"
	"time"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/request"
	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/response"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/commons"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/mappers/users"
)

func (p *Controller) InitAuthRoutes() {
	p.routes["/auth/login"] = p.HandleLogin
}

func (p *Controller) HandleLogin(
	conn net.Conn,
	args json.RawMessage,
	jwtToken *string,
) error {
	var req users.LoginReq
	if err := request.ParseJSON(args, &req); err != nil {
		return err
	}

	ctx := context.Background()
	user, err := p.services.Users().Login(ctx, req.Username, req.Password)
	if err != nil {
		return err
	}

	jwt, err := p.tokenstore.GenerateToken(user.Clients.ID, user.ID, user.Username, user.Nickname, 24*time.Hour)
	if err != nil {
		return err
	}

	return response.Response(commons.StatusOK, "Login successful", jwt)
}
