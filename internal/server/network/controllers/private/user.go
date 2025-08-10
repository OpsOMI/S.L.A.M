package private

import (
	"context"
	"encoding/json"
	"net"

	"github.com/OpsOMI/S.L.A.M/internal/server/domains/commons"
	responseserver "github.com/OpsOMI/S.L.A.M/internal/server/network/response"
	"github.com/OpsOMI/S.L.A.M/internal/server/network/utils"
	"github.com/OpsOMI/S.L.A.M/internal/shared/dto/message"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/response"
)

func (p *Controller) InitUserRoutes() {
	p.routes["/send"] = p.HandleMessage
}

func (p *Controller) HandleMessage(
	conn net.Conn,
	args json.RawMessage,
	jwtToken *string,
) error {
	ctx := context.Background()
	userInfo := p.store.ParseToken(jwtToken)

	var req message.MessageReq
	if err := utils.ParseJSON(args, &req); err != nil {
		p.logger.Error("Failed to parse message request: " + err.Error())
		return nil
	}

	chatRoom, ok := p.connections.GetClientRoom(userInfo.ClientID)
	if !ok {
		p.logger.Warn("User tried to send message without joining a room")
		return responseserver.Response(commons.StatusBadRequest, "Enter a Room/Chat First", nil)
	}

	room, err := p.services.Rooms().GetByCode(ctx, chatRoom)
	if err != nil {
		return err
	}

	connections, ok := p.connections.GetConnectionsByRoomCode(room.Code)
	if !ok {
		p.logger.Info("No active clients in room. Message will be saved to DB.")

		if err := p.services.Messages().CreateMessage(ctx, userInfo.UserID.String(), chatRoom, req.Content); err != nil {
			p.logger.Error("Failed to save room message: " + err.Error())
			return err
		}

		return responseserver.Response(commons.StatusOK, "Message Saved", nil)
	}

	for _, c := range connections {
		if c == conn {
			continue
		}

		payload := message.MessageResp{
			RoomCode:       chatRoom,
			SenderNickname: userInfo.Nickname,
			Content:        req.Content,
		}

		request.Send(c, response.BaseResponse{
			ResponseID: "INCOMING_MESSAGE",
			Code:       commons.StatusOK,
			Message:    "Message Sent!",
			Data:       payload,
		})

	}

	if err := p.services.Messages().CreateMessage(ctx, userInfo.UserID.String(), chatRoom, req.Content); err != nil {
		p.logger.Error("Failed to save room message: " + err.Error())
		return err
	}

	return responseserver.Response(commons.StatusOK, "Message Sent into Room!", nil)
}
