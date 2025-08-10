package private

import (
	"context"
	"encoding/json"
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/server/domains/commons"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/response"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/utils"
	"github.com/OpsOMI/S.L.A.M/internal/shared/dto/rooms"
)

func (p *Controller) InitRoomRoutes() {
	p.routes["/join"] = p.HandleJoin
	p.routes["/room/create"] = p.CreateRoom
	p.routes["/room/list"] = p.List
	p.routes["/room/clean"] = p.Clean
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

	room, err := p.services.Rooms().JoinRoom(ctx, req.RoomCode, req.Password)
	if err != nil {
		return err
	}
	p.connections.SetClientRoom(userInfo.ClientID, room.Code)

	domainMessages, err := p.services.Messages().GetMessagesByRoomCode(ctx, room.Code)
	if err != nil {
		return err
	}

	for i := range domainMessages.Items {
		if domainMessages.Items[i].SenderNickname == userInfo.Nickname {
			domainMessages.Items[i].SenderNickname = "You"
		}
	}

	return response.Response(commons.StatusOK, "Joined Successfully", rooms.OneJoin(req.RoomCode, domainMessages))
}

func (p *Controller) CreateRoom(
	conn net.Conn,
	args json.RawMessage,
	jwtToken *string,
) error {
	ctx := context.Background()
	userInfo := p.store.ParseToken(jwtToken)
	var req rooms.CreateReq
	if err := utils.ParseJSON(args, &req); err != nil {
		return nil
	}

	code, err := p.services.Rooms().Create(ctx, userInfo.UserID.String(), req.Password)
	if err != nil {
		return err
	}

	return response.Response(commons.StatusOK, "Created Successfully", rooms.OneCreate(*code))
}

func (p *Controller) List(
	conn net.Conn,
	args json.RawMessage,
	jwtToken *string,
) error {
	ctx := context.Background()
	userInfo := p.store.ParseToken(jwtToken)
	var req rooms.ListRoomReq
	if err := utils.ParseJSON(args, &req); err != nil {
		return nil
	}

	dbModels, err := p.services.Rooms().GetByOwnerID(ctx, userInfo.UserID.String(), req.Page, req.Limit)
	if err != nil {
		return err
	}

	return response.Response(commons.StatusOK, "Listed Successfully", rooms.ManyRoom(dbModels))
}

func (p *Controller) Clean(
	conn net.Conn,
	args json.RawMessage,
	jwtToken *string,
) error {
	ctx := context.Background()
	userInfo := p.store.ParseToken(jwtToken)
	var req rooms.CleanRoomReq
	if err := utils.ParseJSON(args, &req); err != nil {
		return nil
	}

	if err := p.services.Messages().DeleteMessageInRoom(ctx, userInfo.UserID, req.RoomCode); err != nil {
		return err
	}

	return response.Response(commons.StatusOK, "Messages Deleted Successfully", nil)
}
