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
	"github.com/OpsOMI/S.L.A.M/internal/shared/network/request"
)

func (p *Controller) InitUserRoutes() {
	p.routes["/join"] = p.HandleJoin
	p.routes["/send"] = p.HandleMessage
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
		return response.Response(commons.StatusBadRequest, "Enter a Room/Chat First", nil)
	}

	receiver, isRoom, err := p.services.Rooms().IsIsRoomOrDirectChat(ctx, chatRoom)
	if err != nil {
		p.logger.Error("Failed to determine room or direct chat: " + err.Error())
		return err
	}

	if !isRoom {
		receiverConn, ok := p.connections.GetConn(receiver.ClientKey)
		if !ok {
			p.logger.Info("Receiver client is not online. Message will be saved to DB.")

			// Direct Chat Message
			if err := p.services.Messages().CreateMessage(ctx, userInfo.UserID.String(), receiver.UserID.String(), "", req.Content); err != nil {
				p.logger.Error("Failed to save direct message: " + err.Error())
				return err
			}

			return response.Response(commons.StatusOK, "Message Saved", nil)
		}

		request.Send(receiverConn, message.MessageResp{
			SenderNickname: userInfo.Nickname,
			Content:        req.Content,
		})

		// Direct Chat Message
		if err := p.services.Messages().CreateMessage(ctx, userInfo.UserID.String(), receiver.UserID.String(), "", req.Content); err != nil {
			p.logger.Error("Failed to save direct message: " + err.Error())
			return err
		}

		return response.Response(commons.StatusOK, "Direct Message Sent!", nil)
	}

	connections, ok := p.connections.GetConnectionsByRoomCode(chatRoom)
	if !ok {
		p.logger.Info("No active clients in room. Message will be saved to DB.")

		// Room Message
		if err := p.services.Messages().CreateMessage(ctx, userInfo.UserID.String(), "", chatRoom, req.Content); err != nil {
			p.logger.Error("Failed to save room message: " + err.Error())
			return err
		}

		return response.Response(commons.StatusOK, "Message Saved", nil)
	}

	for _, conn := range connections {
		request.Send(conn, message.MessageResp{
			SenderNickname: userInfo.Nickname,
			Content:        req.Content,
		})
	}

	// Room Message
	if err := p.services.Messages().CreateMessage(ctx, userInfo.UserID.String(), "", chatRoom, req.Content); err != nil {
		p.logger.Error("Failed to save room message: " + err.Error())
		return err
	}

	return response.Response(commons.StatusOK, "Message Sent into Room!", nil)
}
