package owner

import (
	"context"
	"encoding/json"
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/request"
	"github.com/OpsOMI/S.L.A.M/internal/adapters/network/response"
	"github.com/OpsOMI/S.L.A.M/internal/server/domains/commons"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/mappers/users"
)

func (p *Controller) InitAuthRoutes() {
	p.routes["/auth/register"] = p.HandleRegister
}

func (p *Controller) HandleRegister(
	conn net.Conn,
	args json.RawMessage,
	jwtToken *string,
) error {
	var req users.RegisterReq
	if err := request.ParseJSON(args, &req); err != nil {
		return nil
	}

	ctx := context.Background()
	id, err := p.services.Users().CreateUser(ctx, req.Nickname, req.Username, req.Password, "user")
	if err != nil {
		return err
	}

	return response.Response(commons.StatusOK, "User Created!", id)
}
