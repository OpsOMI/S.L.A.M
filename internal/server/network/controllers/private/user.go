package private

import (
	"context"
	"encoding/json"
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/server/domains/commons"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/response"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/utils"
	"github.com/OpsOMI/S.L.A.M/internal/shared/dto/message"
	"github.com/OpsOMI/S.L.A.M/internal/shared/dto/rooms"
)

func (p *Controller) InitUserRoutes() {
	p.routes["/join"] = p.HandleJoin
}

func (p *Controller) HandleJoin(
	conn net.Conn,
	args json.RawMessage,
	jwtToken *string,
) error {
	ctx := context.Background()
	userInfo := p.store.ParseToken(jwtToken)
	var req rooms.JoinReq
	if err := utils.ParseJSON(args, &req); err != nil {
		return nil
	}

	room, err := p.services.Rooms().GetByCode(ctx, req.RoomCode)
	if err != nil {
		return err
	}
	p.connections.SetClientRoom(userInfo.ClientID, room.Code)

	domainMessages, err := p.services.Messages().GetMessagesByRoomCode(ctx, room.Code)
	if err != nil {
		return err
	}

	return response.Response(commons.StatusOK, "Joined Successfully", message.ManyMessage(domainMessages))
}
