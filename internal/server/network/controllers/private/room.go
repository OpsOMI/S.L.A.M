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
	p.routes["/room/create"] = p.CreateRoom
	p.routes["/room/myrooms"] = p.MyRooms
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

	return response.Response(commons.StatusOK, "Joined Successfully", rooms.OneCreate(*code))
}

func (p *Controller) MyRooms(
	conn net.Conn,
	args json.RawMessage,
	jwtToken *string,
) error {
	ctx := context.Background()
	userInfo := p.store.ParseToken(jwtToken)
	var req rooms.MyRoomReq
	if err := utils.ParseJSON(args, &req); err != nil {
		return nil
	}

	dbModels, err := p.services.Rooms().GetByOwnerID(ctx, userInfo.UserID.String(), req.Page, req.Limit)
	if err != nil {
		return err
	}

	return response.Response(commons.StatusOK, "Joined Successfully", rooms.ManyRoom(dbModels))
}
